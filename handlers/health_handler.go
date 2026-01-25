package handlers

import (
	"kasir-api/pkg"
	"net/http"
)

// @Summary Health Check
// @Description Health Check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} pkg.ResponsePayload
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	pkg.ResponseSuccess(w, http.StatusOK, "success", map[string]string{
		"status":  "OK",
		"message": "API Running",
	})
}
