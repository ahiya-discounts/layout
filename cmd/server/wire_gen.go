// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"server/internal/biz"
	"server/internal/conf"
	"server/internal/data"
	"server/internal/dep"
	"server/internal/server"
	"server/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(contextContext context.Context, bootstrap *conf.Bootstrap, confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	textMapPropagator := dep.NewTextMapPropagator()
	tracerProvider, err := dep.NewTracerProvider(contextContext, bootstrap, textMapPropagator)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup, err := data.NewData(confData, logger, tracerProvider)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	usersRepo := data.NewUsersRepo(dataData)
	usersUsecase := biz.NewUsersUsecase(usersRepo, logger)
	usersService := service.NewUsersService(usersUsecase, logger)
	productsRepo := data.NewProductsRepo(dataData)
	productsUsecase := biz.NewProductsUsecase(productsRepo, logger)
	productsService := service.NewProductsService(productsUsecase, logger)
	meterProvider, err := dep.NewMeterProvider(bootstrap)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	meter, err := dep.NewMeter(bootstrap, meterProvider)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	grpcServer, err := server.NewGRPCServer(confServer, greeterService, usersService, productsService, logger, meter, tracerProvider)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewHTTPServer(confServer, greeterService, usersService, productsService, logger, meter, tracerProvider)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
