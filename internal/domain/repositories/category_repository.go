package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/internal/domain/models"
)

type CategoryRepository interface {
	GetAllCategory() ([]models.Category, error)
	CreateCategory(category *models.Category) error
	GetCategoryByID(id int) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id int) error
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (repo *categoryRepository) GetAllCategory() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
func (repo *categoryRepository) CreateCategory(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}
func (repo *categoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"

	var p models.Category
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("Category not found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}
func (repo *categoryRepository) UpdateCategory(category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	err := repo.db.QueryRow(query, category.Name, category.Description, category.ID).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}
func (repo *categoryRepository) DeleteCategory(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	return err
}
