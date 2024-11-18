package api

import (
	"medical-appointments/internal/infrastructure/config"
	"medical-appointments/internal/infrastructure/routes"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// StartServer initializes and starts the server
func InitApplication(deps *config.Dependencies) (err error) {

	deps.Logger.Info("Starting API...")

	// Middlewares
	deps.Router.Use(middleware.Logger)
	//deps.Router.Use(mw.LoggerMiddleware(deps.Logger)) // Probando me falta pulir
	deps.Router.Use(middleware.Recoverer)

	// Register the routes
	routes.Register(deps.Router, deps.Logger, deps.Database)

	// Start HTTP server
	err = http.ListenAndServe(":8080", deps.Router)

	return
}
