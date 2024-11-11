package model

import (
	e "api/internal/shared/error"
)

type TodoProperties struct {
	Id          int64
	Title       string
	Description string
	Completed   bool
}

func (t *TodoProperties) Build() Todo {
	return Todo{t.Id, t.Title, t.Description, t.Completed}
}

type Todo struct {
	id          int64
	title       string
	description string
	completed   bool
}

func (d *Todo) Id() int64 {
	return d.id
}
func (d *Todo) Title() string {
	return d.title
}
func (d *Todo) Description() string {
	return d.description
}
func (d *Todo) Completed() bool {
	return d.completed
}

func TodoBuilder(id int64, title string, description string, completed bool) (*TodoProperties, error) {
	if id == 0 {
		return &TodoProperties{}, e.ErrorBuilder(e.InvalidArgument).Property("id").Build()
	}
	if title == "" {
		return &TodoProperties{}, e.ErrorBuilder(e.InvalidArgument).Property("title").Build()
	}
	return &TodoProperties{id, title, description, completed}, nil
}

func TodoFromRepository(id int64, title string, description string, completed bool) (Todo, error) {
	var model Todo
	if id == 0 {
		return model, e.ErrorBuilder(e.InvalidArgument).Property("id").Build()
	}
	if title == "" {
		return model, e.ErrorBuilder(e.InvalidArgument).Property("title").Build()
	}
	model = Todo{id, title, description, completed}
	return model, nil
}
