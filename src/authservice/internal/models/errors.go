package models

import (
	"errors"
	"net/http"
)

type AuthError struct {
	err  error
	code int
}

func (e AuthError) Error() string {
	return e.err.Error()
}

func (e AuthError) Code() int {
	return e.code
}

var (
	ErrUsernameExists = AuthError{
		err:  errors.New("username already exists"),
		code: http.StatusConflict,
	}

	ErrEmailExists = AuthError{
		err:  errors.New("email already exists"),
		code: http.StatusConflict,
	}
)
