package repository

import (
	"context"
	"social/internal/users/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
