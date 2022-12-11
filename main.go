package main

import (
	"embed"

	"github.com/labstack/echo/v4"

	"github.com/quile8x/healthy_web/config"
	"github.com/quile8x/healthy_web/container"
	"github.com/quile8x/healthy_web/logger"
	"github.com/quile8x/healthy_web/middleware"
	"github.com/quile8x/healthy_web/migration"
	"github.com/quile8x/healthy_web/repository"
	"github.com/quile8x/healthy_web/router"
	"github.com/quile8x/healthy_web/session"
)

//go:embed application.*.yml
var yamlFile embed.FS

//go:embed zaplogger.*.yml
var zapYamlFile embed.FS

//go:embed public/*
var staticFile embed.FS

// @license.name MIT
// @license.url https://opensource.org/licenses/mit-license.php

// @host localhost:8080
// @BasePath /api
func main() {
	e := echo.New()

	conf, env := config.Load(yamlFile)
	logger := logger.InitLogger(env, zapYamlFile)
	logger.GetZapLogger().Infof("Loaded this configuration : application." + env + ".yml")

	rep := repository.NewBookRepository(logger, conf)
	sess := session.NewSession()
	container := container.NewContainer(rep, sess, conf, logger, env)

	migration.CreateDatabase(container)
	migration.InitMasterData(container)

	router.Init(e, container)
	middleware.InitLoggerMiddleware(e, container)
	middleware.InitSessionMiddleware(e, container)
	middleware.StaticContentsMiddleware(e, container, staticFile)

	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
	}

	defer rep.Close()
}
