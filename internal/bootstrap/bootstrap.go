package bootstrap

import (
	"database/sql"
	"kasir-api/internal/config"
	"kasir-api/internal/database"
	"kasir-api/internal/domain/repositories"
	"kasir-api/internal/domain/usecases"
	"kasir-api/internal/http/handlers"
	"kasir-api/internal/http/middleware"
	"kasir-api/internal/pkg"
	"kasir-api/internal/routes"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Bootstrap struct {
	DB     *sql.DB
	Port   string
	Router http.Handler
}

func InitApp() *Bootstrap {
	pkg.InitLogger()
	pkg.Log.Info("running server...")

	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		pkg.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Failed to connect to database")
	}

	deps := initDependencies(db)
	router := routes.RegisterAll(deps)

	return &Bootstrap{
		DB:     db,
		Port:   cfg.Port,
		Router: router,
	}
}

func (a *Bootstrap) Run() {
	addr := "0.0.0.0:" + a.Port

	pkg.Log.WithFields(logrus.Fields{
		"address": addr,
	}).Info("HTTP server running")

	if err := http.ListenAndServe(addr, middleware.Logging(a.Router)); err != nil {
		pkg.Log.Fatal(err)
	}
}

func initDependencies(db *sql.DB) *routes.RouteConfig {
	productRepo := repositories.NewProductRepository(db)
	productUseCase := usecases.NewProductUseCase(productRepo)
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryUseCase := usecases.NewCategoryUseCase(categoryRepo)
	healthUseCase := usecases.NewHealthUseCase("Kasir API", "1.0.0")

	return &routes.RouteConfig{
		ProductHandler:  handlers.NewProductHandler(productUseCase),
		CategoryHandler: handlers.NewCategoryHandler(categoryUseCase),
		HealthHandler:   handlers.NewHealthHandler(healthUseCase),
	}
}
