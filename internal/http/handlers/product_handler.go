package handlers

import (
	"encoding/json"
	"kasir-api/internal/domain/models"
	"kasir-api/internal/domain/usecases"
	"kasir-api/internal/pkg"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	productUseCase usecases.ProductUseCase
}

func NewProductHandler(productUseCase usecases.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: productUseCase}
}

// @Summary Get All Products
// @Description Get All Products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product [get]
func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	pkg.Log.WithFields(logrus.Fields{
		"handler": "product_handler",
		"action":  "get_all_products",
		"method":  r.Method,
	}).Info("Get all products handler called")

	products, err := h.productUseCase.GetAllProducts(name)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "product_handler",
			"action":  "get_all_products",
			"error":   err.Error(),
		}).Error("Failed to get products")
		pkg.ResponseError(w, http.StatusInternalServerError, "Failed to get products", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler": "product_handler",
		"action":  "get_all_products",
		"count":   len(products),
	}).Info("Products retrieved successfully")

	pkg.ResponseSuccess(w, http.StatusOK, "Products retrieved successfully", products)
}

// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body dto.ProductRequest true "Product Request"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	pkg.Log.WithFields(logrus.Fields{
		"handler": "product_handler",
		"action":  "create_product",
		"method":  r.Method,
	}).Info("Create product handler called")

	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "product_handler",
			"action":  "create_product",
			"error":   err.Error(),
		}).Warn("Invalid request body")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	err = h.productUseCase.CreateProduct(&newProduct)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":      "product_handler",
			"action":       "create_product",
			"product_name": newProduct.Name,
			"error":        err.Error(),
		}).Error("Failed to create product")
		pkg.ResponseError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":      "product_handler",
		"action":       "create_product",
		"product_id":   newProduct.ID,
		"product_name": newProduct.Name,
	}).Info("Product created successfully")

	w.WriteHeader(http.StatusCreated)
	pkg.ResponseSuccess(w, http.StatusCreated, "Product created successfully", nil)
}

// @Summary Get Product By ID
// @Description Get Product By ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "product_handler",
			"action":  "get_product_by_id",
			"id_str":  idStr,
		}).Warn("Invalid product ID format")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Product ID", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":    "product_handler",
		"action":     "get_product_by_id",
		"product_id": id,
	}).Info("Get product by ID handler called")

	product, err := h.productUseCase.GetProductByID(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":    "product_handler",
			"action":     "get_product_by_id",
			"product_id": id,
			"error":      err.Error(),
		}).Error("Failed to get product")
		pkg.ResponseError(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":      "product_handler",
		"action":       "get_product_by_id",
		"product_id":   id,
		"product_name": product.Name,
	}).Info("Product found")

	pkg.ResponseSuccess(w, http.StatusOK, "Product found", product)
}

// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.ProductRequest true "Product Request"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "product_handler",
			"action":  "update_product",
			"id_str":  idStr,
		}).Warn("Invalid product ID format")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Product ID", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":    "product_handler",
		"action":     "update_product",
		"product_id": id,
	}).Info("Update product handler called")

	// get data dari request
	var updateProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":    "product_handler",
			"action":     "update_product",
			"product_id": id,
			"error":      err.Error(),
		}).Warn("Invalid request body")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	updateProduct.ID = id
	err = h.productUseCase.UpdateProduct(&updateProduct)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":      "product_handler",
			"action":       "update_product",
			"product_id":   id,
			"product_name": updateProduct.Name,
			"error":        err.Error(),
		}).Error("Failed to update product")
		pkg.ResponseError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":      "product_handler",
		"action":       "update_product",
		"product_id":   id,
		"product_name": updateProduct.Name,
	}).Info("Product updated successfully")

	pkg.ResponseSuccess(w, http.StatusOK, "Product updated successfully", updateProduct)
}

// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "product_handler",
			"action":  "delete_product",
			"id_str":  idStr,
		}).Warn("Invalid product ID format")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Product ID", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":    "product_handler",
		"action":     "delete_product",
		"product_id": id,
	}).Info("Delete product handler called")

	err = h.productUseCase.DeleteProduct(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":    "product_handler",
			"action":     "delete_product",
			"product_id": id,
			"error":      err.Error(),
		}).Error("Failed to delete product")
		pkg.ResponseError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":    "product_handler",
		"action":     "delete_product",
		"product_id": id,
	}).Info("Product deleted successfully")

	pkg.ResponseSuccess(w, http.StatusOK, "Product deleted successfully", nil)
}

// define function untuk handle method
func (h *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	// Debug tracing
	pkg.Log.WithFields(logrus.Fields{
		"handler": "product_handler",
		"func":    "HandleProduct",
		"method":  r.Method,
	}).Info("Routing dispatcher called")

	switch r.Method {
	case http.MethodGet:
		h.GetAllProduct(w, r)
	case http.MethodPost:
		h.CreateProduct(w, r)
	default:
		http.Error(w, "Request not found", http.StatusNotFound)
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	pkg.Log.WithFields(logrus.Fields{
		"handler": "product_handler",
		"func":    "HandleProductByID",
		"method":  r.Method,
		"path":    r.URL.Path,
	}).Info("Routing dispatcher called")

	switch r.Method {
	case http.MethodGet:
		h.GetProductByID(w, r)
	case http.MethodPut:
		h.UpdateProduct(w, r)
	case http.MethodDelete:
		h.DeleteProduct(w, r)
	default:
		http.Error(w, "Request not found", http.StatusMethodNotAllowed)
	}
}
