package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateBudgetRequest struct {
	CategoryID  int64           `json:"category_id" binding:"required"`
	Currency    string          `json:"currency" binding:"required"`
	BudgetLimit decimal.Decimal `json:"budget_limit" binding:"required"`
	PeriodStart time.Time       `json:"period_start" binding:"required"`
	PeriodEnd   time.Time       `json:"period_end" binding:"required"`
}

type UpdateBudgetRequest struct {
	CategoryID  int64           `json:"category_id" binding:"required"`
	Currency    string          `json:"currency" binding:"required"`
	BudgetLimit decimal.Decimal `json:"budget_limit" binding:"required"`
	PeriodStart time.Time       `json:"period_start" binding:"required"`
	PeriodEnd   time.Time       `json:"period_end" binding:"required"`
}

type BudgetResponse struct {
	ID          int64           `json:"id"`
	CategoryID  int64           `json:"category_id"`
	Currency    string          `json:"currency"`
	BudgetLimit decimal.Decimal `json:"budget_limit"`
	PeriodStart time.Time       `json:"period_start"`
	PeriodEnd   time.Time       `json:"period_end"`
}
