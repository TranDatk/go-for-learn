package repository

import (
	"context"
	"database/sql"
	"social/internal/entity"

	"github.com/lib/pq"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) Create(ctx context.Context, p *entity.Post) error {
	query := `
		INSERT INTO posts (id , title, content, user_id, tags)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, title, content, user_id, tags
	`

	err := r.DB.QueryRowContext(ctx, query, p.ID, p.Title, p.Content, p.UserID, pq.Array(p.Tags)).Scan(
		&p.ID, &p.Title, &p.Content, &p.UserID, pq.Array(&p.Tags),
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) (*[]entity.Post, error) {
	query := `
		SELECT id, title, content, user_id, tags
		FROM posts
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entity.Post

	for rows.Next() {
		var p entity.Post

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.UserID,
			pq.Array(&p.Tags),
		)

		if err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return &result, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	var p entity.Post

	err := r.DB.QueryRowContext(ctx,
		"SELECT id, title, content, user_id, tags FROM posts WHERE id=$1", id).
		Scan(&p.ID, &p.Title, &p.Content, &p.UserID, pq.Array(&p.Tags))

	if err != nil {
		return nil, err
	}

	return &p, nil
}
