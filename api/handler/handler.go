package handler

import (	
	"github.com/google/uuid"
)

type UserHandler interface {
	RegisterUser(username, email, password string)
}
