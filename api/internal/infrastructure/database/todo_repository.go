package database

import (
	"api/internal/domain/model"
	"api/internal/infrastructure/database/dto"
	d "api/internal/shared/database"
)

type TodoDatabaseRepository struct {
	tx d.Tx
}

func (t TodoDatabaseRepository) Create(model model.Todo) error {
	tx, err := t.tx.Tx()
	if err != nil {
		return err
	}
	dto := dto.ConvertToTodoDto(model)
	_, err = tx.Exec(`
		INSERT INTO 
			todos (title, description, completed)
		VALUES (?, ?, ?)`, dto.Title, dto.Description, dto.Completed)
	return err
}

func (t TodoDatabaseRepository) Delete(id int64) error {
	tx, err := t.tx.Tx()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		DELETE FROM 
			todos 
		WHERE 
			id = (?)`, id)
	return err
}

func (t TodoDatabaseRepository) Update(model model.Todo) error {
	tx, err := t.tx.Tx()
	if err != nil {
		return err
	}
	dto := dto.ConvertToTodoDto(model)
	_, err = tx.Exec(`
		UPDATE 
			todos 
		SET 
			title = ?, description = ?, completed = ?
		WHERE 
			id = ?`, dto.Title, dto.Description, dto.Completed, dto.Id)
	return err
}

func (t TodoDatabaseRepository) Find(id int64) (model.Todo, error) {
	tx, err := t.tx.Tx()
	if err != nil {
		return model.Todo{}, err
	}
	var dst dto.TodoDto
	err = tx.Select(&dst, `
		SELECT 
			id, 
			title, 
			description, 
			completed 
		FROM 
			todos 
		WHERE 
			id = (?)`, id)
	if err != nil {
		return model.Todo{}, err
	}
	res, err := dto.ConvertToTodoDomain(dst)
	if err != nil {
		return model.Todo{}, err
	}
	return res, nil
}

func NewTodoDatabaseRepository(tx d.Tx) TodoDatabaseRepository {
	return TodoDatabaseRepository{tx}
}
