package entity

import "github.com/google/uuid"

type Uuid struct {
	Id uuid.UUID
}

func (u Uuid) GenerateNewId() string {
	return uuid.New().String()
}
