package service

import (
	"errors"
	"expense-tracker/internal/apperrors"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

type ExpenseRepository interface {
	CreateExpense(userID int64, categoryID int64, amount decimal.Decimal, currency string, expenseDate time.Time, expenseDescription string) (models.Expense, error)
	GetExpenses(userID int64, currency string, categoryID *int64) ([]models.Expense, error)
	GetExpenseByID(expenseID int64, userID int64) (models.Expense, error)
	UpdateExpense(expenseID int64, userID int64, categoryID *int64, amount *decimal.Decimal, currency *string, expenseDate *time.Time, expenseDescription *string) (models.Expense, error)
	DeleteExpenseByID(expenseID int64, userID int64) error
	GetExpenseStats(userID int64) (models.ExpenseStats, error)
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

func (s *ExpenseService) GetExpenses(userID int64, currency string, categoryID *int64) ([]dto.ExpenseResponse, error) {
	expenses, err := s.repo.GetExpenses(userID, currency, categoryID)
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

func (s *ExpenseService) GetExpenseByID(expenseID int64, userID int64) (dto.ExpenseResponse, error) {
	expense, err := s.repo.GetExpenseByID(expenseID, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.ExpenseResponse{}, apperrors.ErrExpenseNotFound
		}

		return dto.ExpenseResponse{}, err
	}

	expenseDTO := dto.ExpenseResponse{
		ID:                 expense.ID,
		CategoryID:         expense.CategoryID,
		Amount:             expense.Amount,
		Currency:           expense.Currency,
		ExpenseDate:        expense.ExpenseDate,
		ExpenseDescription: expense.ExpenseDescription,
		CreatedAt:          expense.CreatedAt,
		UpdatedAt:          expense.UpdatedAt,
	}

	return expenseDTO, nil
}

func (s *ExpenseService) UpdateExpense(expenseID int64, userID int64, categoryID *int64, amount *decimal.Decimal, currency *string, expenseDate *time.Time, expenseDescription *string) (dto.ExpenseResponse, error) {
	expense, err := s.repo.UpdateExpense(expenseID, userID, categoryID, amount, currency, expenseDate, expenseDescription)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.ExpenseResponse{}, apperrors.ErrExpenseNotFound
		}

		return dto.ExpenseResponse{}, err
	}

	expenseDTO := dto.ExpenseResponse{
		ID:                 expense.ID,
		CategoryID:         expense.CategoryID,
		Amount:             expense.Amount,
		Currency:           expense.Currency,
		ExpenseDate:        expense.ExpenseDate,
		ExpenseDescription: expense.ExpenseDescription,
		CreatedAt:          expense.CreatedAt,
		UpdatedAt:          expense.UpdatedAt,
	}

	return expenseDTO, nil
}

func (s *ExpenseService) DeleteExpenseByID(expenseID int64, userID int64) error {
	err := s.repo.DeleteExpenseByID(expenseID, userID)
	if err != nil {
		if errors.Is(err, apperrors.ErrExpenseNotFound) {
			return apperrors.ErrExpenseNotFound
		}

		return err
	}

	return nil
}

func (s *ExpenseService) GetExpenseStats(userID int64) (dto.ExpenseStatsResponse, error) {
	expenseStats, err := s.repo.GetExpenseStats(userID)
	if err != nil {
		return dto.ExpenseStatsResponse{}, err
	}

	expenseStatsDTO := dto.ExpenseStatsResponse{
		TotalAmount:   expenseStats.TotalAmount,
		TotalExpenses: expenseStats.TotalExpenses,
		AverageAmount: expenseStats.AverageAmount,
	}

	return expenseStatsDTO, nil
}
