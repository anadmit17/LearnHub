package repository

import (
	"sync"

	"authservice/internal/models"
)

type InMemoryAuthRepository struct {
	userMap map[string]models.User
	mu      sync.Mutex
}

func NewInMemoryUserRepository() *InMemoryAuthRepository {
	return &InMemoryAuthRepository{
		userMap: make(map[string]models.User),
	}
}

func (r *InMemoryAuthRepository) RegisterUser(user models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.userMap[user.ID] = user

	return nil
}

func (r *InMemoryAuthRepository) Exists(user models.User) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, found := r.userMap[user.ID]

	return found, nil
}
