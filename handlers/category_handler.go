package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/pkg"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Get All Categories
// @Description Get All Categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category [get]
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	pkg.ResponseSuccess(w, http.StatusOK, "success", models.Categories)
}

// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param body body dto.CategoryRequest true "Create Category Request"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category [post]
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Request", nil)
		return
	}

	newCategory.ID = len(models.Categories) + 1
	models.Categories = append(models.Categories, newCategory)

	pkg.ResponseSuccess(w, http.StatusCreated, "success", models.Categories)
}

// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param body body dto.CategoryRequest true "Update Category Request"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category/{id} [put]
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStrCategory := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStrCategory)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Request", nil)
		return
	}

	for i := range models.Categories {
		if models.Categories[i].ID == id {
			models.Categories[i].ID = updateCategory.ID
			models.Categories[i] = updateCategory

			pkg.ResponseSuccess(w, http.StatusOK, "success", models.Categories)
			return
		}
	}

	pkg.ResponseSuccess(w, http.StatusNotFound, "Category not found", nil)
}

// @Summary Get Category By ID
// @Description Get Category By ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category/{id} [get]
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStrCategory := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStrCategory)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	for _, ctg := range models.Categories {
		if ctg.ID == id {
			pkg.ResponseSuccess(w, http.StatusOK, "Success", ctg)
			return
		}
	}

	pkg.ResponseSuccess(w, http.StatusNotFound, "Category not found", nil)
}

// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category/{id} [delete]
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStrCategory := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStrCategory)
	if err != nil {
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	for i, ctg := range models.Categories {
		if ctg.ID == id {
			models.Categories = append(models.Categories[:i], models.Categories[i+1:]...)
			pkg.ResponseSuccess(w, http.StatusOK, "success", nil)
			return
		}
	}

	pkg.ResponseSuccess(w, http.StatusNotFound, "Category not found", nil)
}
