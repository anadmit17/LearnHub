package service

import (
	"authservice/internal/models"
	"authservice/internal/repository"
)

type DefaultAuthService struct {
	repo repository.AuthRepository
}

func NewDefaultAuthService(repo repository.AuthRepository) *DefaultAuthService {
	return &DefaultAuthService{
		repo: repo,
	}
}

func (s *DefaultAuthService) RegisterUser(user models.User) error {
	if exists, _ := s.repo.Exists(user); !exists {
		return models.ErrUserExists
	}

	s.repo.RegisterUser(user)

	return nil
}
