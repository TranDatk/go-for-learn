package service

import (
	"context"
	"fmt"
	"social/internal/entity"
	"social/internal/users/repository"
	"social/pkg/helper"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(r repository.IUserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(ctx context.Context, u *entity.User) error {

	id, err := uuid.NewV7()

	if err != nil {
		return fmt.Errorf("generate uuid: %w", err)
	}

	u.ID = id

	if err := s.repo.Create(ctx, u); err != nil {
		return helper.NewRequestError("repo create user", err)
	}

	return nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return s.repo.GetByID(ctx, id)
}
