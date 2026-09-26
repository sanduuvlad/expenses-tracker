package repository

import (
	"context"
	"expense-tracker/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type BudgetRepository struct {
	pool *pgxpool.Pool
}

func NewBudgetRepository(pool *pgxpool.Pool) *BudgetRepository {
	return &BudgetRepository{
		pool: pool,
	}
}

func (r *BudgetRepository) CreateBudget(userID int64, categoryID int64, currency string, budgetLimit decimal.Decimal, periodStart time.Time, periodEnd time.Time) (models.Budget, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`INSERT INTO budgets (user_id, category_id, currency, budget_limit, period_start, period_end)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, category_id, currency, budget_limit, period_start, period_end`,
		userID,
		categoryID,
		currency,
		budgetLimit,
		periodStart,
		periodEnd,
	)

	var budget models.Budget

	err := row.Scan(
		&budget.ID,
		&budget.UserID,
		&budget.CategoryID,
		&budget.Currency,
		&budget.BudgetLimit,
		&budget.PeriodStart,
		&budget.PeriodEnd,
	)
	if err != nil {
		return models.Budget{}, err
	}

	return budget, nil
}
