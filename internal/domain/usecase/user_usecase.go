package usecase

import (
	"context"

	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/domain"
	"github.com/VadimRight/User_Microserver/internal/domain/dtos"
	"github.com/VadimRight/User_Microserver/internal/domain/entity"
	"github.com/VadimRight/User_Microserver/internal/domain/errors"
)

type userUsecase struct {
	repo domain.Repository
	cfg  config.Config
}

type UserRepositoryRegister interface {
	InsertUser(ctx context.Context, id entity.Uuid, username string, password string) (entity.User, error)
}

type UserRegisterRequestValidator interface {
	ValidateUserRegisterRequest() error
}

func (u *userUsecase) RegisterUser(ctx context.Context, payload dtos.UserRegisterRequest) (userId entity.Uuid, err error) {
	if exists := u.repo.IsUserExist(ctx, payload.Username); exists {
		return userId, errors.Conflict(errors.ErrUsernameAlreadyExists)
	}
	userId.GenerateNewId()
	user, err = u.repo.InsertUser(ctx, userId, payload.Username)
	return userId, nil
}
