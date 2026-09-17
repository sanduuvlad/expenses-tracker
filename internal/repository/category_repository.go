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

func (r *CategoryRepository) GetCategories(userID int64) ([]models.Category, error) {
	rows, err := r.pool.Query(
		context.Background(),
		`SELECT id, user_id, name, created_at
		FROM categories
		WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var category models.Category

		err := rows.Scan(
			&category.ID,
			&category.UserID,
			&category.Name,
			&category.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) UpdateCategory(categoryID int64, userID int64, name string) (models.Category, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`UPDATE categories
		SET name = $1
		WHERE id = $2 AND user_id = $3
		RETURNING id, user_id, name, created_at`,
		name,
		categoryID,
		userID,
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
