package repository

import "authservice/internal/models"

type UserRepository interface {
	CreateUser(user models.User)
	Exists(user models.User) bool
}