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

func (d *TodoProperties) Build() Todo {
	return Todo{d.Id, d.Title, d.Description, d.Completed}
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
		return &TodoProperties{}, e.ErrorBuilder(e.InvalidArgument).Property("token").Build()
	}
	if title == "" {
		return &TodoProperties{}, e.ErrorBuilder(e.InvalidArgument).Property("title").Build()
	}
	if description == "" {
		return &TodoProperties{}, e.ErrorBuilder(e.InvalidArgument).Property("description").Build()
	}
	if completed == false {
		return &TodoProperties{}, e.ErrorBuilder(e.InvalidArgument).Property("completed").Build()
	}
	return &TodoProperties{id, title, description, completed}, nil
}

func TodoFromRepository(id int64, title string, description string, completed bool) (Todo, error) {
	var model Todo
	if id == 0 {
		return model, e.ErrorBuilder(e.InvalidArgument).Property("token").Build()
	}
	if title == "" {
		return model, e.ErrorBuilder(e.InvalidArgument).Property("title").Build()
	}
	if description == "" {
		return model, e.ErrorBuilder(e.InvalidArgument).Property("description").Build()
	}
	if completed == false {
		return model, e.ErrorBuilder(e.InvalidArgument).Property("completed").Build()
	}
	model = Todo{id, title, description, completed}
	return model, nil
}
