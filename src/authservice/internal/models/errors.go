package models

import (
	"fmt"
	"net/http"
)

type AuthError struct {
	Error error
	Code  int
}

var (
	ErrUsernameExists = AuthError{
		Error: fmt.Errorf("User with this username is already exists"),
		Code:  http.StatusConflict,
	}

	ErrEmailExists = AuthError{
		Error: fmt.Errorf("User with this email is already exists"),
		Code:  http.StatusConflict,
	}
)
