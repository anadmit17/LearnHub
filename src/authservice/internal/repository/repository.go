package repository

import "authservice/internal/models"

type AuthRepository interface {
	RegisterUser(user models.User) error
	Exists(user models.User) (bool, error)
}
