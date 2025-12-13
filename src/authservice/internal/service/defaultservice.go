package service

import (
	"authservice/internal/models"
	"authservice/internal/repository"
)

type DefaultUserService struct {
	repo repository.UserRepository
}

func NewDefaultUserService(repo repository.UserRepository) *DefaultUserService {
	return &DefaultUserService{
		repo: repo,
	}
}

func (s *DefaultUserService) CreateUser(user models.User) error {
	if s.repo.Exists(user) {
		return models.ErrUserExists
	}

	s.repo.CreateUser(user)

	return nil
} 