package service

import "authservice/internal/models"

type AuthService interface {
	RegisterUser(user models.User) error
}
