package services

import (
	"crud-api-category/internal/models"
	"crud-api-category/internal/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}
