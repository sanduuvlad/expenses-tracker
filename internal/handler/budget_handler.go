package handler

import (
	"errors"
	"expense-tracker/internal/apperrors"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	service *service.BudgetService
}

func NewBudgetHandler(BudgetService *service.BudgetService) *BudgetHandler {
	return &BudgetHandler{
		service: BudgetService,
	}
}

func (h *BudgetHandler) CreateHandler(c *gin.Context) {
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

	var request dto.CreateBudgetRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	budgetResponseDTO, err := h.service.CreateBudget(userIDInt64, request.CategoryID, request.Currency, request.BudgetLimit, request.PeriodStart, request.PeriodEnd)
	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidBudgetLimit) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget limit"})
			return
		}

		if errors.Is(err, apperrors.ErrInvalidPeriod) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget period"})
			return
		}

		if errors.Is(err, apperrors.ErrInvalidCurrency) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid budget currency"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, budgetResponseDTO)
}
