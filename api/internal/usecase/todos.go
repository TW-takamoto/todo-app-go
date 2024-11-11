package usecase

import (
	"api/internal/domain/model"
	"api/internal/usecase/interfaces"
)

type TodosUsecase struct {
	todoQuery interfaces.TodoQuery
}

func (t TodosUsecase) Execute() ([]model.Todo, error) {
	return t.todoQuery.Get()
}

func NewTodosUsecase(todoQuery interfaces.TodoQuery) TodosUsecase {
	return TodosUsecase{todoQuery: todoQuery}
}
