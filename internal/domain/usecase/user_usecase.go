package usecase

import (
	"context"

	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/domain"
	"github.com/VadimRight/User_Microserver/internal/domain/entity"
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
