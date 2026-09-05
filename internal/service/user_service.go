package service

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
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
