package service

import (
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"
)

type CategoryRepository interface {
	CreateCategory(userID int64, name string) (models.Category, error)
	GetCategories(userID int64) ([]models.Category, error)
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

func (s *CategoryService) GetCategories(userID int64) ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetCategories(userID)
	if err != nil {
		return nil, err
	}

	categoriesResponseDTO := make([]dto.CategoryResponse, 0)

	for _, value := range categories {
		categoryDTO := dto.CategoryResponse{
			ID:        value.ID,
			Name:      value.Name,
			CreatedAt: value.CreatedAt,
		}

		categoriesResponseDTO = append(categoriesResponseDTO, categoryDTO)
	}

	return categoriesResponseDTO, nil
}
