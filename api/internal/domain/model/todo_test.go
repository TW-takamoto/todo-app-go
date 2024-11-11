package model_test

import (
	"testing"

	"api/internal/domain/model"
)

func TestNewTodo(t *testing.T) {
	id := int64(1)
	title := "title"
	description := "description"
	completed := true
	builder, err := model.TodoBuilder(id, title, description, completed)
	if err != nil {
		t.Fatal(err)
	}
	todo := builder.Build()
	if todo.Id() != id {
		t.Error("id is not match")
	}
	if todo.Title() != title {
		t.Error("title is not match")
	}
	if todo.Description() != description {
		t.Error("description is not match")
	}
	if !todo.Completed() {
		t.Error("isActive is not match")
	}
}
