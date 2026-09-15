package handler

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(CategoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: CategoryService,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var categoryRequest dto.CreateCategoryRequest

	err := c.ShouldBindJSON(&categoryRequest)
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

	categoryResponseDTO, err := h.service.CreateCategory(userIDInt64, categoryRequest.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, categoryResponseDTO)
}
