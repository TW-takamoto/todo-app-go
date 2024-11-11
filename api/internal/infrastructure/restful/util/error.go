package util

import (
	"github.com/gin-gonic/gin"
)

type ServerError struct {
	Code    int    `json:"Code"    example:"500"`
	Message string `json:"Message" example:"Lost connection with peer node"`
}

func ErrorHappened(c *gin.Context, status int, err error) {
	e := ServerError{
		Code:    status,
		Message: err.Error(),
	}

	c.IndentedJSON(status, e)
}
