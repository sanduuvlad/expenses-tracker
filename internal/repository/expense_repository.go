package repository

import (
	"context"
	"expense-tracker/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type ExpenseRepository struct {
	pool *pgxpool.Pool
}

func NewExpenseRepository(pool *pgxpool.Pool) *ExpenseRepository {
	return &ExpenseRepository{
		pool: pool,
	}
}

func (r *ExpenseRepository) CreateExpense(userID int64, categoryID int64, amount decimal.Decimal, currency string, expenseDate time.Time, expenseDescription string) (models.Expense, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`INSERT INTO expenses (user_id, category_id, amount, currency, expense_date, expense_description)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, category_id, amount, currency, expense_date, expense_description, created_at, updated_at`,
		userID,
		categoryID,
		amount,
		currency,
		expenseDate,
		expenseDescription,
	)

	var expense models.Expense

	err := row.Scan(
		&expense.ID,
		&expense.UserID,
		&expense.CategoryID,
		&expense.Amount,
		&expense.Currency,
		&expense.ExpenseDate,
		&expense.ExpenseDescription,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	)
	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}
