package interfaces

import (
	"api/internal/domain/model"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	Update(todo *model.Todo) error
	Delete(id int64) error
}
