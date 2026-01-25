package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/pkg"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Get All Products
// @Description Get All Products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product [get]
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	pkg.ResponseSuccess(w, http.StatusOK, "success", models.Products)
}

// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body dto.ProductRequest true "Product Request"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	//masukin data ke dalam variable product
	newProduct.ID = len(models.Products) + 1
	models.Products = append(models.Products, newProduct)

	w.WriteHeader(http.StatusCreated)
	pkg.ResponseSuccess(w, http.StatusCreated, "success", models.Products)
}

// @Summary Get Product By ID
// @Description Get Product By ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product/{id} [get]
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Product ID", nil)
		return
	}

	for _, p := range models.Products {
		if p.ID == id {
			pkg.ResponseSuccess(w, http.StatusOK, "success", p)
			return
		}
	}

	pkg.ResponseSuccess(w, http.StatusNotFound, "Product not found", nil)
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
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Product ID", nil)
		return
	}

	// get data dari request
	var updateProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	// loop product cari id ganti sesuai data dari request
	for i := range models.Products {
		if models.Products[i].ID == id {
			models.Products[i].ID = updateProduct.ID
			models.Products[i] = updateProduct

			pkg.ResponseSuccess(w, http.StatusOK, "success", models.Products)
			return
		}
	}
	pkg.ResponseError(w, http.StatusNotFound, "Product not found", nil)
}

// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/product/{id} [delete]
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Product ID", nil)
		return
	}
	// loop product cari id
	for i, p := range models.Products {
		if p.ID == id {
			// bikin slice baru dengan data sebeluym dan sesudah index
			models.Products = append(models.Products[:i], models.Products[i+1:]...)
			pkg.ResponseSuccess(w, http.StatusOK, "success", nil)
			return
		}
	}
	pkg.ResponseError(w, http.StatusNotFound, "Product not found", nil)
}
