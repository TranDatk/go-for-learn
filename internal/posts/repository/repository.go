package repository

import (
	"context"
	"social/internal/entity"
)

type IPostRepository interface {
	Create(ctx context.Context, p *entity.Post) error
	GetAll(ctx context.Context) (*[]entity.Post, error)
	GetByID(ctx context.Context, id string) (*entity.Post, error)
}
