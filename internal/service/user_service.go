package service

import (
	"errors"
	"expense-tracker/internal/apperrors"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int64) (models.User, error)
	CreateUser(email string, passwordHash string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	usersDTO := make([]dto.UserResponse, 0)

	for _, value := range users {
		userDTO := dto.UserResponse{
			ID:        value.ID,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}

		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO, nil
}

func (s *UserService) GetUserByID(id int64) (dto.UserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	userResponseDTO := dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponseDTO, nil
}

func (s *UserService) RegisterUser(email, password string) (dto.UserResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user, err := s.repo.CreateUser(email, string(passwordHash))
	if err != nil {
		return dto.UserResponse{}, err
	}

	userResponseDTO := dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponseDTO, nil
}

func (s *UserService) LoginUser(email, password string) (dto.UserResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.UserResponse{}, apperrors.ErrInvalidCredentials
		}

		return dto.UserResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return dto.UserResponse{}, apperrors.ErrInvalidCredentials
		}

		return dto.UserResponse{}, err
	}

	userResponse := dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, nil
}
