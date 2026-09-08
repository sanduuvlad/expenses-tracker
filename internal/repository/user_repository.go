package repository

import (
	"context"
	"expense-tracker/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, email, password_hash, created_at, updated_at
		FROM users`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id int64) (models.User, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`SELECT id, email, password_hash, created_at, updated_at
        FROM users
        WHERE id = $1`,
		id,
	)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) CreateUser(email string, passwordHash string) (models.User, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, email, created_at, updated_at`,
		email,
		passwordHash,
	)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`SELECT id, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1`,
		email,
	)

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
