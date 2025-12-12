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
		return fmt.Errorf("generate uuid: %w", err)
	}

	p.ID = id

	if err := s.repo.Create(ctx, p); err != nil {
		return helper.NewRequestError("repo create post", err)
	}

	return nil
}

func (s *PostService) GetAllPost(ctx context.Context) (*[]entity.Post, error) {
	return s.repo.GetAll(ctx)
}

func (s *PostService) GetPostByID(ctx context.Context, id string) (*entity.Post, error) {
	return s.repo.GetByID(ctx, id)
}
