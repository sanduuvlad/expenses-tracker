package service

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
	"time"

	"github.com/shopspring/decimal"
)

type ExpenseRepository interface {
	CreateExpense(userID int64, categoryID int64, amount decimal.Decimal, currency string, expenseDate time.Time, expenseDescription string) (models.Expense, error)
	GetExpenses(userID int64) ([]models.Expense, error)
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

func (s *ExpenseService) GetExpenses(userID int64) ([]dto.ExpenseResponse, error) {
	expenses, err := s.repo.GetExpenses(userID)
	if err != nil {
		return nil, err
	}

	expensesResponseDTO := make([]dto.ExpenseResponse, 0)

	for _, value := range expenses {
		expenseDTO := dto.ExpenseResponse{
			ID:                 value.ID,
			CategoryID:         value.CategoryID,
			Amount:             value.Amount,
			Currency:           value.Currency,
			ExpenseDate:        value.ExpenseDate,
			ExpenseDescription: value.ExpenseDescription,
			CreatedAt:          value.CreatedAt,
			UpdatedAt:          value.UpdatedAt,
		}

		expensesResponseDTO = append(expensesResponseDTO, expenseDTO)
	}

	return expensesResponseDTO, nil
}
