package model

import (
	"github.com/jordan-wright/email"
)

type User struct {
	Username   string      `json:"username"`
	Email      email.Email `json:"email"`
	IsVerified bool        `json:"is_verified"`
	IsActive   bool        `json:"is_active"`
}
