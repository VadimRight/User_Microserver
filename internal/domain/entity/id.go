package entity

import "github.com/google/uuid"

type Uuid struct {
	Id uuid.UUID
}

func (u *Uuid) GenerateNewId() string {
	u.Id = uuid.New()
	return u.Id.String()
}
