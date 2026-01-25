package main

import (
	_ "kasir-api/docs"
	"kasir-api/pkg"
	"kasir-api/routes"
	"net/http"
)

// @title           Kasir API
// @version         1.0
// @description     API untuk sistem kasir sederhana
// @host            localhost:8080
// @BasePath        /
func main() {
	pkg.InitLogger()

	routes.SetupRouter()
	pkg.InfoLogger.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		pkg.ErrorLogger.Println("Failed running server:", err)
	}
}
