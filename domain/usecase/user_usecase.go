package usecase

import (
	"context"

	"github.com/VadimRight/User_Microserver/domain"
	"github.com/VadimRight/User_Microserver/domain/entity"
	"github.com/VadimRight/User_Microserver/internal/config"
)

type userUsecase struct {
	repo domain.Repository
	cfg  config.Config
}

type UserRepositoryRegister interface {
	InsertUser(ctx context.Context, id entity.Uuid, username string, password string) (entity.User, error)
}

func (u *userUsecase) RegisterUser(user UserRepositoryRegister) {
}
