package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    uuid.UUID `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostPayload struct {
	Title   string    `json:"title" validate:"required,min=3,max=200"`
	Content string    `json:"content" validate:"required"`
	UserID  uuid.UUID `json:"user_id" validate:"required,gt=0"`
	Tags    []string  `json:"tags" validate:"omitempty,dive,min=1"`
}
