package service

import (
	"errors"
	"expense-tracker/internal/apperrors"
	"expense-tracker/internal/dto"
	"expense-tracker/internal/models"

	"github.com/jackc/pgx/v5"
)

type CategoryRepository interface {
	CreateCategory(userID int64, name string) (models.Category, error)
	GetCategories(userID int64) ([]models.Category, error)
	UpdateCategory(categoryID int64, userID int64, name string) (models.Category, error)
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

func (s *CategoryService) UpdateCategory(categoryID int64, userID int64, name string) (dto.CategoryResponse, error) {
	category, err := s.repo.UpdateCategory(categoryID, userID, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.CategoryResponse{}, apperrors.ErrCategoryNotFound
		}

		return dto.CategoryResponse{}, err
	}

	categoryResponseDTO := dto.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	return categoryResponseDTO, nil
}
