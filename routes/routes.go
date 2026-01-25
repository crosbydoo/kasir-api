package routes

import (
	"kasir-api/handlers"
	"kasir-api/pkg"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter() {
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetAllProducts(w, r)
		case "POST":
			handlers.CreateProduct(w, r)
		default:
			pkg.ResponseError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetProductByID(w, r)
		case "PUT":
			handlers.UpdateProduct(w, r)
		case "DELETE":
			handlers.DeleteProduct(w, r)
		default:
			pkg.ResponseError(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		}
	})

	// Health check
	http.HandleFunc("/health", handlers.HealthCheck)

	// Swagger
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}
