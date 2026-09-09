package handler

import (
	"errors"
	"expense-tracker/internal/apperrors"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/service"
	"fmt"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, usersDTO)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userDTO, err := h.service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, userDTO)
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userRequest dto.RegisterUserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	userResponseDTO, err := h.service.RegisterUser(
		userRequest.Email,
		userRequest.Password,
	)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, userResponseDTO)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var loginRequest dto.LoginUserRequest

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	userResponseDTO, err := h.service.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, userResponseDTO)
}
