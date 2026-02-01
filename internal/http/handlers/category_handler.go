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

type CategoryHandler struct {
	categoryUseCase usecases.CategoryUseCase
}

func NewCategoryHandler(categoryUseCase usecases.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: categoryUseCase}
}

// @Summary Get All Categories
// @Description Get All Categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category [get]
func (h *CategoryHandler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	pkg.Log.WithFields(logrus.Fields{
		"handler": "category_handler",
		"action":  "get_all_categories",
		"method":  r.Method,
	}).Info("Get all categories handler called")

	categories, err := h.categoryUseCase.GetAllCategory()
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "category_handler",
			"action":  "get_all_categories",
			"error":   err.Error(),
		}).Error("Failed to get categories")
		pkg.ResponseError(w, http.StatusInternalServerError, "Failed to get categories", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler": "category_handler",
		"action":  "get_all_categories",
		"count":   len(categories),
	}).Info("Categories retrieved successfully")

	pkg.ResponseSuccess(w, http.StatusOK, "success", categories)
}

// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param body body dto.CategoryRequest true "Create Category Request"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	pkg.Log.WithFields(logrus.Fields{
		"handler": "category_handler",
		"action":  "create_category",
		"method":  r.Method,
	}).Info("Create category handler called")

	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "category_handler",
			"action":  "create_category",
			"error":   err.Error(),
		}).Warn("Invalid request body")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	err = h.categoryUseCase.CreateCategory(&newCategory)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":       "category_handler",
			"action":        "create_category",
			"category_name": newCategory.Name,
			"error":         err.Error(),
		}).Error("Failed to create category")
		pkg.ResponseError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":       "category_handler",
		"action":        "create_category",
		"category_id":   newCategory.ID,
		"category_name": newCategory.Name,
	}).Info("Category created successfully")

	w.WriteHeader(http.StatusCreated)
	pkg.ResponseSuccess(w, http.StatusCreated, "Category created successfully", nil)
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
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	//get id from request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "category_handler",
			"action":  "update_category",
			"id_str":  idStr,
		}).Warn("Invalid category ID format")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":     "category_handler",
		"action":      "update_category",
		"category_id": id,
	}).Info("Update category handler called")

	// get data dari request
	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":     "category_handler",
			"action":      "update_category",
			"category_id": id,
			"error":       err.Error(),
		}).Warn("Invalid request body")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	updateCategory.ID = id
	err = h.categoryUseCase.UpdateCategory(&updateCategory)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":       "category_handler",
			"action":        "update_category",
			"category_id":   id,
			"category_name": updateCategory.Name,
			"error":         err.Error(),
		}).Error("Failed to update category")
		pkg.ResponseError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":       "category_handler",
		"action":        "update_category",
		"category_id":   id,
		"category_name": updateCategory.Name,
	}).Info("Category updated successfully")

	pkg.ResponseSuccess(w, http.StatusOK, "Category updated successfully", updateCategory)
}

// @Summary Get Category By ID
// @Description Get Category By ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category/{id} [get]
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":     "category_handler",
			"action":      "get_category_by_id",
			"category_id": id,
		}).Warn("Invalid category ID format")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":     "category_handler",
		"action":      "get_category_by_id",
		"category_id": id,
	}).Info("Get category by ID handler called")

	category, err := h.categoryUseCase.GetCategoryByID(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":     "category_handler",
			"action":      "get_category_by_id",
			"category_id": id,
			"error":       err.Error(),
		}).Error("Failed to get category")
		pkg.ResponseError(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":       "category_handler",
		"action":        "get_category_by_id",
		"category_id":   id,
		"category_name": category.Name,
	}).Info("Category found")

	pkg.ResponseSuccess(w, http.StatusOK, "Category found", category)
}

// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} pkg.ResponsePayload
// @Router /api/category/{id} [delete]
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	// ganti jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "category_handler",
			"action":  "delete_category",
			"id_str":  idStr,
		}).Warn("Invalid category ID format")
		pkg.ResponseError(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":     "category_handler",
		"action":      "delete_category",
		"category_id": id,
	}).Info("Delete category handler called")

	err = h.categoryUseCase.DeleteCategory(id)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler":     "category_handler",
			"action":      "delete_category",
			"category_id": id,
			"error":       err.Error(),
		}).Error("Failed to delete category")
		pkg.ResponseError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler":     "category_handler",
		"action":      "delete_category",
		"category_id": id,
	}).Info("Category deleted successfully")

	pkg.ResponseSuccess(w, http.StatusOK, "Category deleted successfully", nil)
}

func (h *CategoryHandler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	// Debug tracing
	pkg.Log.WithFields(logrus.Fields{
		"handler": "category_handler",
		"func":    "HandleCategory",
		"method":  r.Method,
	}).Info("Routing dispatcher called")

	switch r.Method {
	case http.MethodGet:
		h.GetAllCategory(w, r)
	case http.MethodPost:
		h.CreateCategory(w, r)
	default:
		http.Error(w, "Request not found", http.StatusNotFound)
	}
}

func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
	pkg.Log.WithFields(logrus.Fields{
		"handler": "category_handler",
		"func":    "GetAllCategory",
		"method":  r.Method,
		"path":    r.URL.Path,
	}).Info("Routing dispatcher called")

	switch r.Method {
	case http.MethodGet:
		h.GetCategoryByID(w, r)
	case http.MethodPut:
		h.UpdateCategory(w, r)
	case http.MethodDelete:
		h.DeleteCategory(w, r)
	default:
		http.Error(w, "Request not found", http.StatusMethodNotAllowed)
	}
}
