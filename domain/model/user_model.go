package model

import (
	"github.com/jordan-wright/email"
)

type User struct {
	Username   string
	Email      email.Email
	IsVerified bool
	IsActive   bool
}
