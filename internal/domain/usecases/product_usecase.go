package usecases

import (
	"errors"
	"kasir-api/internal/domain/models"
	"kasir-api/internal/domain/repositories"
	"kasir-api/internal/pkg"

	"github.com/sirupsen/logrus"
)

// ProductUseCase adalah interface untuk product use cases
type ProductUseCase interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product *models.Product) error
	GetProductByID(id int) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id int) error
}

type productUseCase struct {
	productRepo repositories.ProductRepository
}

// NewProductUseCase membuat instance baru dari ProductUseCase sebagai contructor menerima interface
func NewProductUseCase(productRepo repositories.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepo: productRepo,
	}
}

// GetAllProducts mengambil semua produk dari database
func (uc *productUseCase) GetAllProducts() ([]models.Product, error) {
	pkg.Log.WithFields(logrus.Fields{
		"usecase": "product",
		"action":  "get_all_products",
	}).Info("Executing get all products use case")

	products, err := uc.productRepo.GetAllProduct()
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "product",
			"action":  "get_all_products",
			"error":   err.Error(),
		}).Error("Failed to get all products")
		return nil, err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase": "product",
		"action":  "get_all_products",
		"count":   len(products),
	}).Info("Successfully retrieved all products")

	return products, nil
}

// GetProductByID mengambil produk berdasarkan ID
func (uc *productUseCase) GetProductByID(id int) (*models.Product, error) {
	pkg.Log.WithFields(logrus.Fields{
		"usecase":    "product",
		"action":     "get_product_by_id",
		"product_id": id,
	}).Info("Executing get product by ID use case")

	if id <= 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "get_product_by_id",
			"product_id": id,
		}).Warn("Invalid product ID")
		return nil, errors.New("invalid product ID")
	}

	product, err := uc.productRepo.GetProductByID(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "get_product_by_id",
			"product_id": id,
			"error":      err.Error(),
		}).Error("Failed to get product by ID")
		return nil, err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase":      "product",
		"action":       "get_product_by_id",
		"product_id":   id,
		"product_name": product.Name,
	}).Info("Successfully retrieved product")

	return product, nil
}

// CreateProduct membuat produk baru dengan validasi
func (uc *productUseCase) CreateProduct(product *models.Product) error {
	pkg.Log.WithFields(logrus.Fields{
		"usecase":      "product",
		"action":       "create_product",
		"product_name": product.Name,
	}).Info("Executing create product use case")

	// Validasi business rules
	if product.Name == "" {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "product",
			"action":  "create_product",
		}).Warn("Product name is required")
		return errors.New("product name is required")
	}

	if product.Price < 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "product",
			"action":  "create_product",
			"price":   product.Price,
		}).Warn("Product price cannot be negative")
		return errors.New("product price cannot be negative")
	}

	if product.Stock < 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase": "product",
			"action":  "create_product",
			"stock":   product.Stock,
		}).Warn("Product stock cannot be negative")
		return errors.New("product stock cannot be negative")
	}

	if product.CategoryID <= 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":     "product",
			"action":      "create_product",
			"category_id": product.CategoryID,
		}).Warn("Product category ID is required")
		return errors.New("product category ID is required")
	}

	err := uc.productRepo.CreateProduct(product)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":      "product",
			"action":       "create_product",
			"product_name": product.Name,
			"error":        err.Error(),
		}).Error("Failed to create product")
		return err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase":      "product",
		"action":       "create_product",
		"product_id":   product.ID,
		"product_name": product.Name,
	}).Info("Successfully created product")

	return nil
}

// UpdateProduct mengupdate produk dengan validasi
func (uc *productUseCase) UpdateProduct(product *models.Product) error {
	pkg.Log.WithFields(logrus.Fields{
		"usecase":      "product",
		"action":       "update_product",
		"product_id":   product.ID,
		"product_name": product.Name,
	}).Info("Executing update product use case")

	// Validasi business rules
	if product.ID <= 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "update_product",
			"product_id": product.ID,
		}).Warn("Invalid product ID")
		return errors.New("invalid product ID")
	}

	if product.Name == "" {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "update_product",
			"product_id": product.ID,
		}).Warn("Product name is required")
		return errors.New("product name is required")
	}

	if product.Price < 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "update_product",
			"product_id": product.ID,
			"price":      product.Price,
		}).Warn("Product price cannot be negative")
		return errors.New("product price cannot be negative")
	}

	if product.Stock < 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "update_product",
			"product_id": product.ID,
			"stock":      product.Stock,
		}).Warn("Product stock cannot be negative")
		return errors.New("product stock cannot be negative")
	}

	if product.CategoryID <= 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":     "product",
			"action":      "update_product",
			"product_id":  product.ID,
			"category_id": product.CategoryID,
		}).Warn("Product category ID is required")
		return errors.New("product category ID is required")
	}

	// Cek apakah produk ada
	existingProduct, err := uc.productRepo.GetProductByID(product.ID)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "update_product",
			"product_id": product.ID,
			"error":      err.Error(),
		}).Error("Product not found")
		return errors.New("product not found")
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase":    "product",
		"action":     "update_product",
		"product_id": product.ID,
		"old_name":   existingProduct.Name,
		"new_name":   product.Name,
	}).Info("Updating product")

	err = uc.productRepo.UpdateProduct(product)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":      "product",
			"action":       "update_product",
			"product_id":   product.ID,
			"product_name": product.Name,
			"error":        err.Error(),
		}).Error("Failed to update product")
		return err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase":      "product",
		"action":       "update_product",
		"product_id":   product.ID,
		"product_name": product.Name,
	}).Info("Successfully updated product")

	return nil
}

// DeleteProduct menghapus produk berdasarkan ID
func (uc *productUseCase) DeleteProduct(id int) error {
	pkg.Log.WithFields(logrus.Fields{
		"usecase":    "product",
		"action":     "delete_product",
		"product_id": id,
	}).Info("Executing delete product use case")

	if id <= 0 {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "delete_product",
			"product_id": id,
		}).Warn("Invalid product ID")
		return errors.New("invalid product ID")
	}

	// Cek apakah produk ada sebelum dihapus
	product, err := uc.productRepo.GetProductByID(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "delete_product",
			"product_id": id,
			"error":      err.Error(),
		}).Error("Product not found")
		return errors.New("product not found")
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase":      "product",
		"action":       "delete_product",
		"product_id":   id,
		"product_name": product.Name,
	}).Info("Deleting product")

	err = uc.productRepo.DeleteProduct(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"usecase":    "product",
			"action":     "delete_product",
			"product_id": id,
			"error":      err.Error(),
		}).Error("Failed to delete product")
		return err
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase":      "product",
		"action":       "delete_product",
		"product_id":   id,
		"product_name": product.Name,
	}).Info("Successfully deleted product")

	return nil
}
