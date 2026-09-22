package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateExpenseRequest struct {
	CategoryID         int64           `json:"category_id" binding:"required"`
	Amount             decimal.Decimal `json:"amount" binding:"required"`
	Currency           string          `json:"currency" binding:"required"`
	ExpenseDate        time.Time       `json:"expense_date" binding:"required"`
	ExpenseDescription string          `json:"expense_description" binding:"required"`
}

type ExpenseResponse struct {
	ID                 int64           `json:"id"`
	CategoryID         int64           `json:"category_id"`
	Amount             decimal.Decimal `json:"amount"`
	Currency           string          `json:"currency"`
	ExpenseDate        time.Time       `json:"expense_date"`
	ExpenseDescription string          `json:"expense_description"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

type UpdateExpenseRequest struct {
	CategoryID         *int64           `json:"category_id"`
	Amount             *decimal.Decimal `json:"amount"`
	Currency           *string          `json:"currency"`
	ExpenseDate        *time.Time       `json:"expense_date"`
	ExpenseDescription *string          `json:"expense_description"`
}

type ExpenseStatsResponse struct {
	TotalAmount   decimal.Decimal `json:"total_amount"`
	TotalExpenses int64           `json:"total_expenses"`
	AverageAmount decimal.Decimal `json:"average_amount"`
}
