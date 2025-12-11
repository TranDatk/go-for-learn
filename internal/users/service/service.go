package service

import (
	"context"
	"social/internal/users/entity"
	"social/internal/users/repository"
)

type UserService struct {
	repo repository.IUserRepository
}

type IUserService interface {
	Register(ctx context.Context, u *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
