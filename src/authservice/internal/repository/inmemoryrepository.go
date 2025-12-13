package repository

import (
	"sync"

	"authservice/internal/models"
)

type InMemoryUserRepository struct {
	userMap map[string]models.User
	mu      sync.Mutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		userMap: make(map[string]models.User),
	}
}

func (r *InMemoryUserRepository) CreateUser(user models.User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.userMap[user.ID] = user
}

func (r *InMemoryUserRepository) Exists(user models.User) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, found := r.userMap[user.ID]
	
	return found
}
