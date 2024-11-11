package restful

import (
	"api/internal/interface_adapter/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/todos", func(c *gin.Context) {
		controller.TodosController(c)
	})
}
