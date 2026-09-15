package service

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
)

type CategoryRepository interface {
	CreateCategory(userID int64, name string) (models.Category, error)
}

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) CreateCategory(userID int64, name string) (dto.CategoryResponse, error) {
	category, err := s.repo.CreateCategory(userID, name)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	categoryDTO := dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	return categoryDTO, nil
}
