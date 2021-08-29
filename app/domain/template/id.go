package template

import (
	"github.com/google/uuid"
)

type Id struct {
	value uuid.UUID
}

func (i Id) Value() uuid.UUID {
	return i.value
}

func NewId(value uuid.UUID) Id {
	res := new(Id)
	res.value = value
	return *res
}

func ParseId(value string) (Id, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return Id{}, err
	}

	res := new(Id)
	res.value = id

	return *res, nil
}

func (Id) Create() Id {
	res := new(Id)
	id, _ := uuid.NewRandom()
	res.value = id
	return *res
}
