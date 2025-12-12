package repository

import (
	"context"
	"database/sql"
	"social/internal/entity"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u *entity.User) error {
	_, err := r.DB.ExecContext(ctx,
		"INSERT INTO users (id, name, username, password) VALUES ($1, $2, $3, $4)", u.ID, u.Name, u.Username, u.Password)

	return err
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var u entity.User

	err := r.DB.QueryRowContext(ctx,
		"SELECT id, name, username FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.Name, &u.Username)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
