package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	ID                 int64
	UserID             int64
	CategoryID         int64
	Amount             decimal.Decimal
	Currency           string
	ExpenseDate        time.Time
	ExpenseDescription string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type ExpenseStats struct {
	TotalAmount   decimal.Decimal
	TotalExpenses int64
	AverageAmount decimal.Decimal
}
