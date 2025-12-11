package repository

import (
	"context"
	"errors"
	"social/internal/users/entity"
	"sync"
)

type MemoryRepository struct {
	mu    sync.RWMutex
	store map[string]*entity.User
}

func NewMemory() *MemoryRepository {
	return &MemoryRepository{
		store: make(map[string]*entity.User),
	}
}

func (m *MemoryRepository) Create(ctx context.Context, u *entity.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[u.ID.String()] = u
	return nil
}

func (m *MemoryRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	user, ok := m.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}
