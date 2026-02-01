package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/internal/domain/models"
)

// Interface digunakan sebagai kontrak
type ProductRepository interface {
	GetAllProduct() ([]models.Product, error)
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

func (repo *productRepository) GetAllProduct() ([]models.Product, error) {
	query := "SELECT id, name, price, stock FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (repo productRepository) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID)

	return err
}

func (repo *productRepository) GetProductByID(id int) (*models.Product, error) {
	query := "SELECT id, name, price, stock FROM products WHERE id = $1"

	var p models.Product
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("Product not found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo productRepository) UpdateProduct(product *models.Product) error {
	query := "UPDATE products SET name = $2, price = $3, stock = $4 WHERE id = $1"
	_, err := repo.db.Exec(query, product.ID, product.Name, product.Price, product.Stock)
	return err
}

func (repo productRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	return err
}
