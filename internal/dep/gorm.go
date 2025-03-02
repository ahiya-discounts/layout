package dep

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gtracing "gorm.io/plugin/opentelemetry/tracing"
	"server/internal/conf"
)

type Gorm struct {
	// TODO wrapped database client
	db     *gorm.DB
	logger log.Logger
}

func openDB(c *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DriverName:           c.Database.Driver,
				DSN:                  c.Database.Source,
				PreferSimpleProtocol: true,
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, errors.InternalServer("failed to open DB", err.Error())
	}
	return db, nil
}

func NewGorm(c *conf.Data, logger log.Logger, tp trace.TracerProvider) (*Gorm, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db, err := openDB(c)
	if err != nil {
		return nil, nil, err
	}

	if err := db.Use(gtracing.NewPlugin(
		gtracing.WithTracerProvider(tp),
	)); err != nil {
		panic(err)
	}

	return &Gorm{
		db:     db,
		logger: logger,
	}, cleanup, nil
	return nil, nil, nil
}

func GormMigrate(ctx context.Context, c *conf.Data, logger log.Logger, models ...interface{}) {
	_, span := otel.Tracer("data").Start(ctx, "Migrate")
	defer span.End()

	log.NewHelper(logger).Info("migrating the schema")
	client, err := openDB(c)
	if err != nil {
		log.NewHelper(logger).Error("failed opening database: %v", err)
	}
	err = client.AutoMigrate(models...)
	if err != nil {
		log.NewHelper(logger).Error("failed migrating the schema: %v", err)
	}
}
