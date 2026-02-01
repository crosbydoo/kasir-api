package usecases

import (
	"kasir-api/internal/domain/models"
	"kasir-api/internal/pkg"
	"time"

	"github.com/sirupsen/logrus"
)

// HealthUseCase adalah interface untuk health check use case
type HealthUseCase interface {
	CheckHealth() (*models.HealthResponse, error)
}

type healthUseCase struct {
	serviceName string
	version     string
}

// NewHealthUseCase membuat instance baru dari HealthUseCase
func NewHealthUseCase(serviceName, version string) HealthUseCase {
	return &healthUseCase{
		serviceName: serviceName,
		version:     version,
	}
}

// CheckHealth melakukan pengecekan kesehatan aplikasi
func (h *healthUseCase) CheckHealth() (*models.HealthResponse, error) {
	pkg.Log.WithFields(logrus.Fields{
		"usecase": "health_check",
		"action":  "check_health",
	}).Info("Executing health check use case")

	response := &models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   h.serviceName,
		Version:   h.version,
	}

	pkg.Log.WithFields(logrus.Fields{
		"usecase": "health_check",
		"action":  "check_health",
		"status":  response.Status,
	}).Info("Health check completed successfully")

	return response, nil
}
