package usecase

import (
	"api/internal/domain/model"
	"api/internal/usecase/interfaces"
)

type TodoUsecase struct {
	todoQuery interfaces.TodoQuery
}

func (d TodoUsecase) Execute(request interfaces.TodoDetailRequest) (model.Todo, error) {
	return d.todoQuery.Detail(request)
}

func NewTodoUsecase(todoQuery interfaces.TodoQuery) TodosUsecase {
	return TodosUsecase{todoQuery: todoQuery}
}
