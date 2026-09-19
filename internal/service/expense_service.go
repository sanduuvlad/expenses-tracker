package service

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
	"time"

	"github.com/shopspring/decimal"
)

type ExpenseRepository interface {
	CreateExpense(userID int64, categoryID int64, amount decimal.Decimal, currency string, expenseDate time.Time, expenseDescription string) (models.Expense, error)
}

type ExpenseService struct {
	repo ExpenseRepository
}

func NewExpenseService(repo ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) CreateExpense(userID int64, categoryID int64, amount decimal.Decimal, currency string, expenseDate time.Time, expenseDescription string) (dto.ExpenseResponse, error) {
	expense, err := s.repo.CreateExpense(userID, categoryID, amount, currency, expenseDate, expenseDescription)
	if err != nil {
		return dto.ExpenseResponse{}, err
	}

	expenseResponseDTO := dto.ExpenseResponse{
		ID:                 expense.ID,
		CategoryID:         expense.CategoryID,
		Amount:             expense.Amount,
		Currency:           expense.Currency,
		ExpenseDate:        expense.ExpenseDate,
		ExpenseDescription: expense.ExpenseDescription,
		CreatedAt:          expense.CreatedAt,
		UpdatedAt:          expense.UpdatedAt,
	}

	return expenseResponseDTO, nil
}
