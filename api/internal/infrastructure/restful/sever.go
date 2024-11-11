package restful

import (
	"api/internal/infrastructure/restful/middleware"
	"api/internal/shared/config"
	d "api/internal/shared/database"
	"api/internal/shared/logger"

	"github.com/gin-gonic/gin"
)

func Start(config config.Config) error {
	logger, close, err := logger.Setup()
	if err != nil {
		return err
	}
	defer close()
	db, close, err := d.Pool(config)
	if err != nil {
		return err
	}
	defer close()
	r := gin.New()
	//docs.SwaggerInfo.BasePath = "/"
	r.Use(gin.Recovery())
	r.Use(middleware.ConfigMiddleware(config))
	r.Use(middleware.LogMiddleware(logger))
	r.Use(middleware.DatabaseMiddleware(db))
	SetupRoutes(r)
	r.Run()
	return nil
}
