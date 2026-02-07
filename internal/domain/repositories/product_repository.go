package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/internal/domain/models"
)

// Interface digunakan sebagai kontrak
type ProductRepository interface {
	GetAllProduct(name string) ([]models.Product, error)
	GetProductByID(id int) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int) error
}

// CONCRETE IMPLEMENTATION
type productRepository struct {
	db *sql.DB
}

// Constructor RETURN INTERFACE
func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (repo *productRepository) GetAllProduct(name string) ([]models.Product, error) {
	query := "SELECT id, name, price, stock, category_id FROM products"
	args := []interface{}{}
	if name != "" {
		query += " WHERE name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (repo productRepository) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)

	return err
}

func (repo *productRepository) GetProductByID(id int) (*models.Product, error) {
	query := "SELECT p.id, p.name, p.price, p.stock, p.category_id, c.id,  c.name, c.description FROM products p JOIN categories c ON c.id = p.category_id WHERE p.id = $1"

	var p models.Product
	p.Category = &models.Category{}
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.Category.ID, &p.Category.Name, &p.Category.Description)
	if err == sql.ErrNoRows {
		return nil, errors.New("Product not found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo productRepository) UpdateProduct(product *models.Product) error {
	query := "UPDATE products SET name = $2, price = $3, stock = $4, category_id = $5 WHERE id = $1"
	_, err := repo.db.Exec(query, product.ID, product.Name, product.Price, product.Stock, product.CategoryID)
	return err
}

func (repo productRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	return err
}
