package controller

import (
	"net/http"

	"api/internal/infrastructure/database"
	"api/internal/infrastructure/restful/middleware"
	"api/internal/infrastructure/restful/util"
	"api/internal/usecase"

	"github.com/gin-gonic/gin"
)

// Todos godoc
// @Summary タスク・一覧
// @Schemes
// @Description タスクの一覧を返す
// @Tags todos
// @Accept json
// @Produce json
// @Param refs query []string false "refで検索する"
// @Success 200 {object} []Todo
// @Failure	400	{object} util.ServerError
// @Failure	500	{object} util.ServerError
// @Router /todos [get]
func TodosController(c *gin.Context) {
	pool, err := middleware.GetPool(c)
	if err != nil {
		util.ErrorHappened(c, http.StatusInternalServerError, err)
		return
	}
	todoQuery := database.NewTodoDatabaseQuery(pool)
	todos, err := usecase.NewTodosUsecase(todoQuery).Execute()
	if err != nil {
		util.ErrorHappened(c, http.StatusInternalServerError, err)
		return
	}
	res := make([]Todo, len(todos))
	for i, t := range todos {
		res[i] = ConvertToTodoResponse(t)
	}
	c.IndentedJSON(http.StatusOK, res)
}
