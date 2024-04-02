// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/jassue/gin-wire/app/command"
	command2 "github.com/jassue/gin-wire/app/command/handler"
	"github.com/jassue/gin-wire/app/compo"
	"github.com/jassue/gin-wire/app/cron"
	"github.com/jassue/gin-wire/app/data"
	"github.com/jassue/gin-wire/app/handler/app"
	"github.com/jassue/gin-wire/app/handler/common"
	"github.com/jassue/gin-wire/app/middleware"
	"github.com/jassue/gin-wire/app/service"
	"github.com/jassue/gin-wire/config"
	"github.com/jassue/gin-wire/router"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Injectors from wire.go:

// wireApp init application.
func wireApp(configuration *config.Configuration, lumberjackLogger *lumberjack.Logger, zapLogger *zap.Logger) (*App, func(), error) {
	db := data.NewDB(configuration, zapLogger)
	client := data.NewRedis(configuration, zapLogger)
	sonyflake := compo.NewSonyFlake()
	database := data.NewMongoDB(configuration, zapLogger)
	dataData, cleanup, err := data.NewData(zapLogger, db, client, sonyflake, database)
	if err != nil {
		return nil, nil, err
	}
	jwtRepo := data.NewJwtRepo(dataData, zapLogger)
	userRepo := data.NewUserRepo(dataData, zapLogger)
	transaction := data.NewTransaction(dataData)
	userService := service.NewUserService(userRepo, transaction)
	lockBuilder := compo.NewLockBuilder(client)
	jwtService := service.NewJwtService(configuration, zapLogger, jwtRepo, userService, lockBuilder)
	jwtAuth := middleware.NewJWTAuthM(configuration, jwtService)
	recovery := middleware.NewRecoveryM(lumberjackLogger)
	cors := middleware.NewCorsM()
	limiterManager := compo.NewLimiterManager()
	limiter := middleware.NewLimiterM(limiterManager)
	authHandler := app.NewAuthHandler(zapLogger, jwtService, userService)
	storage := compo.NewStorage(configuration, zapLogger)
	mediaRepo := data.NewMediaRepo(dataData, zapLogger, storage)
	mediaService := service.NewMediaService(configuration, zapLogger, mediaRepo, storage)
	uploadHandler := common.NewUploadHandler(zapLogger, mediaService)
	compRepo := data.NewCompRepo(dataData, zapLogger)
	caMsgRepo := data.NewCAMsgRepo(dataData, zapLogger)
	compService := service.NewCompService(compRepo, transaction, caMsgRepo)
	compHandler := app.NewCompHandler(zapLogger, jwtService, compService)
	pageRepo := data.NewPageRepo(dataData, zapLogger)
	pageService := service.NewPageService(pageRepo, transaction, caMsgRepo)
	pageHandler := app.NewPageHandler(zapLogger, jwtService, pageService)
	engine := router.NewRouter(configuration, jwtAuth, recovery, cors, limiter, authHandler, uploadHandler, compHandler, pageHandler)
	server := newHttpServer(configuration, engine)
	exampleJob := cron.NewExampleJob(zapLogger)
	cronCron := cron.NewCron(dataData, zapLogger, exampleJob)
	mainApp := newApp(configuration, zapLogger, server, cronCron)
	return mainApp, func() {
		cleanup()
	}, nil
}

// wireCommand init application.
func wireCommand(configuration *config.Configuration, lumberjackLogger *lumberjack.Logger, zapLogger *zap.Logger) (*command.Command, func(), error) {
	exampleHandler := command2.NewExampleHandler(zapLogger)
	db := data.NewDB(configuration, zapLogger)
	migrateHandler := command2.NewMigrateHandler(zapLogger, db)
	commandCommand := command.NewCommand(exampleHandler, migrateHandler)
	return commandCommand, func() {
	}, nil
}
