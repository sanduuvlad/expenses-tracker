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

func (r *BudgetRepository) GetAllBudgets(userID int64) ([]models.Budget, error) {
	rows, err := r.pool.Query(
		context.Background(),
		`SELECT id, user_id, category_id, currency,
			budget_limit, period_start, period_end
		FROM budgets
		WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	var budgets []models.Budget

	for rows.Next() {
		var budget models.Budget

		err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.CategoryID,
			&budget.Currency,
			&budget.BudgetLimit,
			&budget.PeriodStart,
			&budget.PeriodEnd,
		)
		if err != nil {
			return nil, err
		}

		budgets = append(budgets, budget)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return budgets, nil
}

func (r *BudgetRepository) GetBudgetByID(budgetID int64, userID int64) (models.Budget, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`SELECT id, user_id, category_id, currency,
			budget_limit, period_start, period_end
		FROM budgets
		WHERE id = $1 AND user_id = $2`,
		budgetID,
		userID,
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

func (r *BudgetRepository) UpdateBudget(budgetID int64, userID int64, categoryID int64, currency string, budgetLimit decimal.Decimal, periodStart time.Time, periodEnd time.Time) (models.Budget, error) {
	row := r.pool.QueryRow(
		context.Background(),
		`UPDATE budgets
		SET category_id = $1, currency = $2, budget_limit = $3, 
			period_start = $4, period_end = $5
		WHERE id = $6 AND user_id = $7
		RETURNING id, user_id, category_id, currency, budget_limit, period_start, period_end`,
		categoryID,
		currency,
		budgetLimit,
		periodStart,
		periodEnd,
		budgetID,
		userID,
	)

	var budgetUpdate models.Budget

	err := row.Scan(
		&budgetUpdate.ID,
		&budgetUpdate.UserID,
		&budgetUpdate.CategoryID,
		&budgetUpdate.Currency,
		&budgetUpdate.BudgetLimit,
		&budgetUpdate.PeriodStart,
		&budgetUpdate.PeriodEnd,
	)
	if err != nil {
		return models.Budget{}, err
	}

	return budgetUpdate, nil

}
