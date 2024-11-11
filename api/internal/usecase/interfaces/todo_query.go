package interfaces

import (
	"api/internal/domain/model"
)

type TodoDetailRequest struct {
	Id int64
}

type TodoQuery interface {
	Get() ([]model.Todo, error)
	Detail(request TodoDetailRequest) (model.Todo, error)
}
