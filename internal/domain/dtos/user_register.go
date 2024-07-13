package dtos

import "github.com/invopop/validation"

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterResponce struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (cup UserRegisterRequest) ValidateUserRegisterRequest() error {
	return validation.ValidateStruct(&cup,
		validation.Field(&cup.Username, validation.Required),
		validation.Field(&cup.Email, validation.Required),
		validation.Field(&cup.Password, validation.Required),
	)
}
