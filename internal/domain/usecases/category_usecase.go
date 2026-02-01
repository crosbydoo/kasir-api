package usecases

import (
	"errors"
	"kasir-api/internal/domain/models"
	"kasir-api/internal/domain/repositories"
	"kasir-api/internal/pkg"

	"github.com/sirupsen/logrus"
)

type CategoryUseCase interface {
	GetAllCategory() ([]models.Category, error)
	CreateCategory(category *models.Category) error
	GetCategoryByID(id int) (*models.Category, error)
	UpdateCategory(*models.Category) error
	DeleteCategory(id int) error
}

type categoryUseCase struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repositories.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *categoryUseCase) GetAllCategory() ([]models.Category, error) {
	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "get_all_category",
	}).Info("Executing get all category use case")

	categories, err := uc.categoryRepo.GetAllCategory()

	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "get_all_category",
		}).Error("Failed to get all category use case")
		return nil, err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "get_all_category",
		"count":   len(categories),
	}).Info("Successfully retrieved all categories")

	return categories, nil
}

func (uc *categoryUseCase) CreateCategory(category *models.Category) error {
	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "create_category",
	}).Info("Executing create category use case")

	if category.Name == "" {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "create_category",
		}).Warn("Category name is required")
		return errors.New("category name is required")
	}

	if category.Description == "" {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "create_category",
		}).Warn("Category description is required")
		return errors.New("category description is required")
	}

	return uc.categoryRepo.CreateCategory(category)
}

func (uc *categoryUseCase) GetCategoryByID(id int) (*models.Category, error) {
	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "get_category_by_id",
		"id":      id,
	}).Info("Executing get category by ID use case")

	category, err := uc.categoryRepo.GetCategoryByID(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "get_category_by_id",
			"id":      id,
		}).Error("Failed to get category by ID use case")
		return nil, err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "get_category_by_id",
		"id":      id,
	}).Info("Successfully retrieved category")

	return category, nil
}

func (uc *categoryUseCase) UpdateCategory(category *models.Category) error {
	if category.ID <= 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "update_category",
			"id":      category.ID,
		}).Warn("Invalid category ID")
		return errors.New("invalid category ID")
	}

	if category.Name == "" {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "update_category",
			"id":      category.ID,
		}).Warn("Category name is required")
		return errors.New("category name is required")
	}

	if category.Description == "" {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "update_category",
			"id":      category.ID,
		}).Warn("Category description is required")
		return errors.New("category description is required")
	}

	existingCategory, err := uc.categoryRepo.GetCategoryByID(category.ID)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "update_category",
			"id":      category.ID,
			"error":   err.Error(),
		}).Error("Category not found")
		return errors.New("category not found")
	}
	pkg.Log.WithFields(logrus.Fields{
		"usecase":     "category",
		"action":      "update_category",
		"category_id": category.ID,
		"old_name":    existingCategory.Name,
		"new_name":    category.Name,
	}).Info("Updating category")

	err = uc.categoryRepo.UpdateCategory(category)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "update_category",
			"id":      category.ID,
			"error":   err.Error(),
		}).Error("Failed to update category")
		return err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "update_category",
		"id":      category.ID,
	}).Info("Executing update category use case")

	return nil
}

func (uc *categoryUseCase) DeleteCategory(id int) error {
	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "delete_category",
		"id":      id,
	}).Info("Executing delete category use case")

	if id <= 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "delete_category",
			"id":      id,
		}).Warn("Invalid category ID")
		return errors.New("invalid category ID")
	}

	existingCategory, err := uc.categoryRepo.GetCategoryByID(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "delete_category",
			"id":      id,
			"error":   err.Error(),
		}).Error("Category not found")
		return errors.New("category not found")
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase":       "category",
		"action":        "delete_category",
		"category_id":   id,
		"category_name": existingCategory.Name,
	}).Info("Deleting category")

	err = uc.categoryRepo.DeleteCategory(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "category",
			"action":  "delete_category",
			"id":      id,
			"error":   err.Error(),
		}).Error("Failed to delete category")
		return err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase": "category",
		"action":  "delete_category",
		"id":      id,
	}).Info("Executing delete category use case")

	return nil
}
