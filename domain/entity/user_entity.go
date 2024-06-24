package entity

import (
	"context"
)

type User struct {
	Id         string
	Username   string
	Email      string
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
	return u.Id
}

func (u User) GetName() string {
	return u.Username
}

func (u User) GetEmail() string {
	return u.Email
}
