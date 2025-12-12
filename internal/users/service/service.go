package service

import (
	"context"
	"social/internal/entity"
)

type IUserService interface {
	Register(ctx context.Context, u *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
