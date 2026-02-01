package models

import "time"

// HealthResponse adalah response untuk health check
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}
