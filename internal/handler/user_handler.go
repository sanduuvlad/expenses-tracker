package handler

import (
	"expense-tracker/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{
		service: UserService,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	usersDTO, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, usersDTO)
}
