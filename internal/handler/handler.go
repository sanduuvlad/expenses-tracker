package handler

import (
	"expense-tracker/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.UserService
}

func NewHandler(UserService *service.UserService) *Handler {
	return &Handler{
		service: UserService,
	}
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	usersDTO, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, usersDTO)
}
