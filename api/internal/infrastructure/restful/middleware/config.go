package middleware

import (
	"api/internal/shared/config"
	e "api/internal/shared/error"

	"github.com/gin-gonic/gin"
)

func ConfigMiddleware(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

func GetConfig(c *gin.Context) (config.Config, error) {
	conf, ok := c.Get("config")
	if !ok {
		return config.Config{}, e.ErrorBuilder(e.NotFound).Build()
	}
	conf_, ok := conf.(config.Config)
	if !ok {
		return config.Config{}, e.ErrorBuilder(e.Unknown).Build()
	}
	return conf_, nil
}
