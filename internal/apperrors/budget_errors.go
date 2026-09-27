package apperrors

import "errors"

var ErrInvalidBudgetLimit = errors.New("invalid budget limit")

var ErrInvalidPeriod = errors.New("invalid budget period")

var ErrInvalidCurrency = errors.New("invalid budget currency")

var ErrBudgetNotFound = errors.New("budget not found")
