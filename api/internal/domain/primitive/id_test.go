package primitive_test

import (
	"testing"

	"api/internal/domain/primitive"
)

func TestNewId(t *testing.T) {
	_, err := primitive.NewID()
	if err != nil {
		t.Error(err)

	}
}

func TestFromString(t *testing.T) {
	target := "1537b7ae-641a-45ef-a42d-49cfe7456328"
	id, err := primitive.IDFromString(target)
	if err != nil {
		t.Error(err)

	}
	if id.Value() != target {
		t.Error("id is not match")
	}
}
