package service

import (
	"context"
	"social/internal/entity"
)

type IPostService interface {
	CreatePost(ctx context.Context, p *entity.Post) error
	GetAllPost(ctx context.Context) (*[]entity.Post, error)
	GetPostByID(ctx context.Context, id string) (*entity.Post, error)
}
