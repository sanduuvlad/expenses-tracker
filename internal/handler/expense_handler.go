package handler

import (
	"errors"
	"expense-tracker/internal/apperrors"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	service *service.ExpenseService
}

func NewExpenseHandler(ExpenseService *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: ExpenseService,
	}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	var request dto.CreateExpenseRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	expenseResponseDTO, err := h.service.CreateExpense(userIDInt64, request.CategoryID, request.Amount, request.Currency, request.ExpenseDate, request.ExpenseDescription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, expenseResponseDTO)
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	currencyRequest := c.Query("currency")
	categoryID := c.Query("category_id")

	var categoryIDInt64 *int64

	if categoryID != "" {
		parsedCategoryID, err := strconv.ParseInt(categoryID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		categoryIDInt64 = &parsedCategoryID
	}

	expensesResponse, err := h.service.GetExpenses(userIDInt64, currencyRequest, categoryIDInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, expensesResponse)
}

func (h *ExpenseHandler) GetExpenseByID(c *gin.Context) {
	idStr := c.Param("id")

	expenseID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	expenseResponseDTO, err := h.service.GetExpenseByID(expenseID, userIDInt64)
	if err != nil {
		if errors.Is(err, apperrors.ErrExpenseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, expenseResponseDTO)
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	idStr := c.Param("id")

	expenseID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	var request dto.UpdateExpenseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	expenseResponseDTO, err := h.service.UpdateExpense(expenseID, userIDInt64, request.CategoryID,
		request.Amount, request.Currency, request.ExpenseDate, request.ExpenseDescription)
	if err != nil {
		if errors.Is(err, apperrors.ErrExpenseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, expenseResponseDTO)
}

func (h *ExpenseHandler) DeleteExpenseByID(c *gin.Context) {
	idStr := c.Param("id")

	expenseID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	err = h.service.DeleteExpenseByID(expenseID, userIDInt64)
	if err != nil {
		if errors.Is(err, apperrors.ErrExpenseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ExpenseHandler) GetExpenseStats(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userIDInt64, ok := userID.(int64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	expenseStatsResponse, err := h.service.GetExpenseStats(userIDInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, expenseStatsResponse)
}
