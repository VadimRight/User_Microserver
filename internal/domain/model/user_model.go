package model

import (
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

type User struct {
	Username   string
	Email      email.Email
	IsVerified bool
	IsActive   bool
}

type LoginUser struct {
	Username string
	Email    email.Email
	Password string
}

type GetUser struct {
	Id         uuid.UUID
	Username   string
	Email      string
	IsVerified bool
	IsActive   bool
}
