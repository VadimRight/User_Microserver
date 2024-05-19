package responce

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Responce struct {
	Status string `json:"status"`
	Error string `json:"error omitempty"`
}

const (
	StatusOK = "OK"
	StatusError = "Error"
)

func OK() Responce {
	return Responce {
		Status: StatusOK,
	}
}

func Error(msg string) Responce {
	return Responce {
		Status: StatusError,
		Error: msg,
	}
}

func ErrorValidation(errs validator.ValidationErrors) {
	
}
