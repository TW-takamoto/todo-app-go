package dto

import (
	"api/internal/domain/model"
	e "api/internal/shared/error"
	"time"
)

type TodoDto struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ConvertToTodoDomain(dto TodoDto) (model.Todo, error) {
	res, err := model.TodoFromRepository(dto.Id, dto.Title, dto.Description, dto.Completed)
	if err != nil {
		return model.Todo{}, e.ErrorBuilder(e.NotFound).Build()
	}
	return res, nil
}

func ConvertToTodoDto(domain model.Todo) TodoDto {
	return TodoDto{Id: domain.Id(), Title: domain.Title(), Description: domain.Description(), Completed: domain.Completed()}
}
