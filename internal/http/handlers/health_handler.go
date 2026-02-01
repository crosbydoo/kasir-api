package handlers

import (
	"kasir-api/internal/domain/usecases"
	"kasir-api/internal/pkg"
	"net/http"

	"github.com/sirupsen/logrus"
)

// HealthHandler menangani HTTP request untuk health check
type HealthHandler struct {
	healthUseCase usecases.HealthUseCase
}

// NewHealthHandler membuat instance baru dari HealthHandler
func NewHealthHandler(healthUseCase usecases.HealthUseCase) *HealthHandler {
	return &HealthHandler{
		healthUseCase: healthUseCase,
	}
}

// CheckHealth adalah handler untuk endpoint GET /api/health
// @Summary Health Check
// @Description Mengecek kesehatan aplikasi
// @Tags Health
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /api/health [get]
func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	pkg.Log.WithFields(logrus.Fields{
		"handler": "health_handler",
		"action":  "check_health",
		"method":  r.Method,
	}).Info("Health check handler called")

	// Panggil use case
	response, err := h.healthUseCase.CheckHealth()
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"handler": "health_handler",
			"action":  "check_health",
			"error":   err.Error(),
		}).Error("Failed to check health")

		pkg.Log.Error(w, http.StatusInternalServerError, "Failed to check health")
		return
	}

	pkg.Log.WithFields(logrus.Fields{
		"handler": "health_handler",
		"action":  "check_health",
		"status":  response.Status,
	}).Info("Health check handler completed")

	// Kirim response
	pkg.ResponseSuccess(w, http.StatusOK, "Health check successful", response)
}
