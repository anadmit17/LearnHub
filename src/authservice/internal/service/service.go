package service

import "authservice/internal/models"

type UserService interface {
	CreateUser(user models.User) error
}