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
	ErrUserExists = AuthError{
		err:  errors.New("user already exists"),
		code: http.StatusConflict,
	}
)
