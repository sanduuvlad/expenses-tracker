package service

import (
	"expense-tracker/internal/apperrors"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
	"time"

	"github.com/shopspring/decimal"
)

var allowedCurrencies = map[string]struct{}{
	"MDL": {},
	"EUR": {},
	"USD": {},
}

type BudgetRepository interface {
	CreateBudget(userID int64, categoryID int64, currency string, budgetLimit decimal.Decimal, periodStart time.Time, periodEnd time.Time) (models.Budget, error)
}

type BudgetService struct {
	repo BudgetRepository
}

func NewBudgetService(budgetService BudgetRepository) *BudgetService {
	return &BudgetService{
		repo: budgetService,
	}
}

func (s *BudgetService) CreateBudget(userID int64, categoryID int64, currency string, budgetLimit decimal.Decimal, periodStart time.Time, periodEnd time.Time) (dto.BudgetResponse, error) {
	if budgetLimit.LessThanOrEqual(decimal.Zero) {
		return dto.BudgetResponse{}, apperrors.ErrInvalidBudgetLimit
	}

	if periodStart.After(periodEnd) {
		return dto.BudgetResponse{}, apperrors.ErrInvalidPeriod
	}

	_, ok := allowedCurrencies[currency]
	if !ok {
		return dto.BudgetResponse{}, apperrors.ErrInvalidCurrency
	}

	budgetResponse, err := s.repo.CreateBudget(userID, categoryID, currency, budgetLimit, periodStart, periodEnd)
	if err != nil {
		return dto.BudgetResponse{}, err
	}

	budgetResponseDTO := dto.BudgetResponse{
		ID:          budgetResponse.ID,
		CategoryID:  budgetResponse.CategoryID,
		Currency:    budgetResponse.Currency,
		BudgetLimit: budgetResponse.BudgetLimit,
		PeriodStart: budgetResponse.PeriodStart,
		PeriodEnd:   budgetResponse.PeriodEnd,
	}

	return budgetResponseDTO, nil
}
