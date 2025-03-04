package data

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"server/internal/conf"
	"server/internal/dep"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUsersRepo, NewProductsRepo)

// Data .
type Data struct {
	gorm   *gorm.DB
	mongo  *mongo.Database
	logger log.Logger
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, tp trace.TracerProvider) (*Data, func(), error) {
	g, _, err := dep.NewGorm(c, logger, tp)
	if err != nil {
		return nil, nil, err
	}

	m, clean, err := dep.NewMongo(c, logger)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		clean()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		gorm:   g.DB,
		mongo:  m.DB,
		logger: logger,
	}, cleanup, nil
}
