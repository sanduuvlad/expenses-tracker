package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Budget struct {
	ID          int64
	UserID      int64
	CategoryID  int64
	Currency    string
	BudgetLimit decimal.Decimal
	PeriodStart time.Time
	PeriodEnd   time.Time
}
