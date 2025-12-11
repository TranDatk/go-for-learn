package repository

import (
	"context"
	"database/sql"
	"social/internal/users/entity"
)

type SQLRepository struct {
	DB *sql.DB
}

func NewSQL(db *sql.DB) *SQLRepository {
	return &SQLRepository{DB: db}
}

func (r *SQLRepository) Create(ctx context.Context, u *entity.User) error {
	_, err := r.DB.ExecContext(ctx,
		"INSERT INTO users (name) VALUES ($1)", u.Name)

	return err
}

func (r *SQLRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var u entity.User

	err := r.DB.QueryRowContext(ctx,
		"SELECT id, name FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.Name)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
