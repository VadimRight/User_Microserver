package domain

import (
	"context"

	"github.com/VadimRight/User_Microserver/domain/entity"
)

type Repository interface {
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	InsertUser(ctx context.Context, id entity.Uuid, username string, password string)
	GetUserByID(ctx context.Context, userID string) (entity.User, error)
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	IsUserExist(ctx context.Context, email string) bool
}
