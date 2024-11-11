package controller

import (
	"api/internal/domain/model"
	"time"
)

type Todo struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

func ConvertToTodoResponse(t model.Todo) Todo {
	return Todo{
		Id:          t.Id(),
		Title:       t.Title(),
		Description: t.Description(),
		Completed:   t.Completed(),
	}
}
