package domain

import (
	"context"

	"github.com/VadimRight/User_Microserver/domain/entity"
)

type UserSelecter interface {
	SelectUserByUsername(ctx context.Context, username string) (entity.User, error)
	SelectAllUsers(ctx context.Context) ([]entity.User, error)
	SelectUserByID(ctx context.Context, userID string) (entity.User, error)
}

type UserInserter interface {
	InsertUser(ctx context.Context, id string, username string, password string) (entity.User, error)
}
