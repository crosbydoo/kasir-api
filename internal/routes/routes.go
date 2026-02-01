package routes

import (
	"kasir-api/internal/http/handlers"
	"net/http"
)

type RouteConfig struct {
	ProductHandler  *handlers.ProductHandler
	CategoryHandler *handlers.CategoryHandler
	HealthHandler   *handlers.HealthHandler
}

func RegisterAll(cfg *RouteConfig) http.Handler {
	mux := http.NewServeMux()

	// health
	mux.Handle("/api/health", http.HandlerFunc(cfg.HealthHandler.CheckHealth))

	// product collection
	mux.Handle("/api/product", http.HandlerFunc(cfg.ProductHandler.HandleProduct))
	// product by id
	mux.Handle("/api/product/", http.HandlerFunc(cfg.ProductHandler.HandleProductByID))

	mux.Handle("/api/category", http.HandlerFunc(cfg.CategoryHandler.HandleCategory))
	mux.Handle("/api/category/", http.HandlerFunc(cfg.CategoryHandler.HandleCategoryByID))

	return mux
}
