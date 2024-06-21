package entity

import (
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

type User struct {
	Id         uuid.UUID
	Username   string
	Email      email.Email
	Password   string
	IsVerified bool
	IsActive   bool
}

type User interface {
}
