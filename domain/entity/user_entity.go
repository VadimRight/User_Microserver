package entity

import (
	"context"

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

type UserCreater interface {
	UserCreate(ctx context.Context, username string, password string) (*User, error)
}

type UserGeter interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByID(ctx context.Context, userID string) (User, error)
}

func (u User) GetId() string {
	return u.Id.String()
}

func (u User) GetName() string {
	return u.Username
}

func (u User) GetEmail() email.Email {
	return u.Email
}
