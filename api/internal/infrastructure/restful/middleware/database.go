package middleware

import (
	e "api/internal/shared/error"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func DatabaseMiddleware(pool *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("pool", pool)
		c.Next()
	}
}

func GetPool(c *gin.Context) (*sqlx.DB, error) {
	pool, ok := c.Get("pool")
	if !ok {
		return nil, e.ErrorBuilder(e.NotFound).Build()
	}
	pool_, ok := pool.(*sqlx.DB)
	if !ok {
		return nil, e.ErrorBuilder(e.Unknown).Build()
	}
	return pool_, nil
}
