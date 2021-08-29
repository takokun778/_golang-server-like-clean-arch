package template

import (
	"time"
)

type Time struct {
	value time.Time
}

func (t Time) Value() time.Time {
	return t.value
}

func NewTime(value time.Time) Time {
	res := new(Time)
	res.value = value.UTC()
	return *res
}

func (Time) Now() Time {
	res := new(Time)
	res.value = time.Now().UTC()
	return *res
}
