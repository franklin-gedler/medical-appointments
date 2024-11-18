package routes

import (
	"medical-appointments/internal/application/controllers"
	"medical-appointments/internal/domain/ping"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Register registra las rutas de la aplicación
func Register(r *chi.Mux, logger *zap.Logger, db *gorm.DB) {
	// Crear el repositorio y el servicio
	pingRepo := ping.NewPingRepository(db)
	pingService := ping.NewPingService(pingRepo)
	pingController := controllers.NewPingController(pingService)

	// Versionar la API
	r.Route("/api/v1", func(r chi.Router) {
		// Endpoint de ping
		r.Route("/ping", func(r chi.Router) {
			r.Get("/", pingController.Ping(logger))
		})

		// 	Example login and register routes
		//	r.Route("/auth", func(r chi.Router) {
		//		r.Post("/login", authController.Login)
		//		r.Post("/register", authController.Register)
		//	})

		// Example users routes, CRUD operations
		//	r.Route("/users", func(r chi.Router) {
		//		r.Use(middleware.AuthMiddleware) // Añade el middleware de autenticación de token para todo el grupo de rutas
		//		r.Get("/", userController.GetUsers)
		//		r.Post("/", userController.CreateUser)
		//		r.Get("/{id}", userController.GetUser)
		//		r.Put("/{id}", userController.UpdateUser)
		//		r.Delete("/{id}", userController.DeleteUser)
		//	})

	})
}
