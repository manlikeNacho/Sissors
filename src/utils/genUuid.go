package utils

import (
	"github.com/google/uuid"
)

type UuidGenerator interface {
	IdGenerate() string
}

type uuidGenr struct{}

var GenerateUUID UuidGenerator = &uuidGenr{}

func (u *uuidGenr) IdGenerate() string {
	Uuid, _ := uuid.NewRandom()
	id := Uuid.String()
	return id
}
