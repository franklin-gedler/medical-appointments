package config

import (
	"medical-appointments/internal/infrastructure/database"
	"medical-appointments/internal/infrastructure/logger"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Dependencies is the dependencies struct
type Dependencies struct {
	// Router is the router for the API
	Router *chi.Mux
	// Database is the database connection
	Database *gorm.DB
	// Logger for injection
	Logger *zap.Logger
}

func InitializeDependencies() (*Dependencies, error) {

	// Initialize logger
	loggerConfig := logger.LoggerConfig{UseJSON: false} // logger with format JSON change to true
	logCtx := logger.InitializeLoggerContext(loggerConfig)
	logInstance := logger.FromContext(logCtx)

	godotenv.Load("./.env")

	// New router
	router := chi.NewRouter()

	// Database connection
	db := database.NewDatabaseInstance()
	dbConnection, err := db.Connect()
	if err != nil {
		return nil, err
	}
	logInstance.Info("Connected to database")

	logInstance.Info("Dependencies initialized")

	return &Dependencies{
		Router:   router,
		Database: dbConnection,
		Logger:   logInstance,
	}, nil
}
