package repository

import (
	"context"
	"expense-tracker/internal/models"
	"fmt"
	"strings"
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

func (r *ExpenseRepository) GetExpenses(userID int64) ([]models.Expense, error) {
	rows, err := r.pool.Query(
		context.Background(),
		`SELECT id, user_id, category_id, amount, currency,
			expense_date, expense_description, created_at, updated_at
		FROM expenses
		WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense

	for rows.Next() {
		var expense models.Expense

		err := rows.Scan(
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
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpenseRepository) GetExpenseByID(expenseID int64, userID int64) (models.Expense, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`SELECT id, user_id, category_id, amount, currency,
			expense_date, expense_description, created_at, updated_at
		FROM expenses
		WHERE id = $1 AND user_id = $2`,
		expenseID,
		userID,
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

func (r *ExpenseRepository) UpdateExpense(expenseID int64, userID int64, categoryID *int64, amount *decimal.Decimal, currency *string, expenseDate *time.Time, expenseDescription *string) (models.Expense, error) {
	var setParts []string

	var args []any

	if categoryID != nil {
		parameterNumber := len(args) + 1

		setParts = append(setParts, fmt.Sprintf("category_id = $%d", parameterNumber))
		args = append(args, *categoryID)
	}

	if amount != nil {
		parameterNumber := len(args) + 1

		setParts = append(setParts, fmt.Sprintf("amount = $%d", parameterNumber))
		args = append(args, *amount)
	}

	if currency != nil {
		parameterNumber := len(args) + 1

		setParts = append(setParts, fmt.Sprintf("currency = $%d", parameterNumber))
		args = append(args, *currency)
	}

	if expenseDate != nil {
		parameterNumber := len(args) + 1

		setParts = append(setParts, fmt.Sprintf("expense_date = $%d", parameterNumber))
		args = append(args, *expenseDate)
	}

	if expenseDescription != nil {
		parameterNumber := len(args) + 1

		setParts = append(setParts, fmt.Sprintf("expense_description = $%d", parameterNumber))
		args = append(args, *expenseDescription)
	}

	setParts = append(setParts, "updated_at = NOW()")

	expenseIDParameter := len(args) + 1
	args = append(args, expenseID)

	userIDParameter := len(args) + 1
	args = append(args, userID)

	setClause := strings.Join(setParts, ", ")

	query := fmt.Sprintf(`
		UPDATE expenses
		SET %s
		WHERE id = $%d
			AND user_id = $%d
		RETURNING id, user_id, category_id, amount, currency,
			expense_date, expense_description, created_at, updated_at`,
		setClause,
		expenseIDParameter,
		userIDParameter,
	)

	row := r.pool.QueryRow(
		context.Background(),
		query,
		args...,
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
