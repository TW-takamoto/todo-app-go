package primitive

import (
	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func (i *ID) Value() string {
	return i.value.String()
}

func NewID() (ID, error) {
	value, err := uuid.NewRandom()
	return ID{value: value}, err
}

func IDFromString(target string) (ID, error) {
	value, err := uuid.Parse(target)
	return ID{value}, err
}
