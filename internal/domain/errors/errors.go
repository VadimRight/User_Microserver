package errors

import "errors"

var (
	ErrUsernameAlreadyExists = errors.New("user with this username already exists")
	ErrEmailAlreadyExists    = errors.New("user with this email already exists")
)
