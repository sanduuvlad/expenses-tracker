package repository

import (
	"context"
	"expense-tracker/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository struct {
	pool *pgxpool.Pool
}

func NewCategoryRepository(pool *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{
		pool: pool,
	}
}

func (r *CategoryRepository) CreateCategory(userID int64, name string) (models.Category, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`INSERT INTO categories (user_id, name)
		VALUES ($1, $2)
		RETURNING id, user_id, name, created_at`,
		userID,
		name,
	)

	var category models.Category

	err := row.Scan(
		&category.ID,
		&category.UserID,
		&category.Name,
		&category.CreatedAt,
	)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}
