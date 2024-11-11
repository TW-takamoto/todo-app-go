package database

import (
	"api/internal/domain/model"
	"api/internal/infrastructure/database/dto"
	"api/internal/usecase/interfaces"

	"github.com/jmoiron/sqlx"
)

type TodoDatabaseQuery struct {
	pool *sqlx.DB
}

func (t TodoDatabaseQuery) Get() ([]model.Todo, error) {
	var res []model.Todo
	var dst []dto.TodoDto
	err := t.pool.Select(&dst, "SELECT id, title, description, completed, created_at, updated_at FROM todos")
	if err != nil {
		return nil, err
	}

	//結果を変換
	for _, t := range dst {
		domain, err := dto.ConvertToTodoDomain(t)
		if err != nil {
			continue
		}
		res = append(res, domain)
	}

	return res, nil
}

func (d TodoDatabaseQuery) Detail(request interfaces.TodoDetailRequest) (model.Todo, error) {
	var res model.Todo
	var dst dto.TodoDto
	err := d.pool.Get(&dst, "SELECT id, title, description, completed FROM todos WHERE id = (?)", request.Id)
	if err != nil {
		return res, err
	}

	return dto.ConvertToTodoDomain(dst)
}

func NewTodoDatabaseQuery(pool *sqlx.DB) TodoDatabaseQuery {
	return TodoDatabaseQuery{pool: pool}
}
