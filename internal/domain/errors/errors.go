package errors

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code    int
	Err     error
	Message string
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

var (
	ErrUsernameAlreadyExists = errors.New("user with this username already exists")
	ErrEmailAlreadyExists    = errors.New("user with this email already exists")
)

func Conflict(err error) error {
	return &AppError{
		Code:    http.StatusConflict,
		Message: "Conflict",
		Err:     err,
	}
}
