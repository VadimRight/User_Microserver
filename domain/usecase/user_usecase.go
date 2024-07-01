package usecase

import (
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/repository"
)

type userUsecase struct {
	repo repository.UserRepository
	cfg  config.Config
}

func (u *userUsecase) Register() {
}
