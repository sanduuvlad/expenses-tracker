package handler

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/service"
	"net/http"

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

	expensesResponse, err := h.service.GetExpenses(userIDInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, expensesResponse)
}
