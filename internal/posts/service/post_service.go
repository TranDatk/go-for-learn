package service

import (
	"context"
	"fmt"
	"social/internal/entity"
	"social/internal/posts/repository"
	"social/pkg/helper"

	"github.com/google/uuid"
)

type PostService struct {
	repo repository.IPostRepository
}

func NewPostService(r repository.IPostRepository) *PostService {
	return &PostService{repo: r}
}

func (s *PostService) CreatePost(ctx context.Context, p *entity.Post) error {
	id, err := uuid.NewV7()

	if err != nil {
		return fmt.Errorf("Generate uuid: %w", err)
	}

	p.ID = id

	if err := s.repo.Create(ctx, p); err != nil {
		return helper.NewCustomError(err, "Create post failed", nil)
	}

	return nil
}

func (s *PostService) GetAllPost(ctx context.Context) (*[]entity.Post, error) {
	posts, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, helper.NewCustomError(err, "Get all post failed", nil)
	}
	return posts, nil
}

func (s *PostService) GetPostByID(ctx context.Context, id string) (*entity.Post, error) {
	post, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, helper.NewCustomError(err, "Get post failed", nil)
	}

	return post, nil
}
