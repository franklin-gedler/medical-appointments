package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase es una interfaz que define el método Connect
type NewDatabase interface {
	Connect() (*gorm.DB, error)
}

// Database es una estructura que contiene la conexión a la base de datos
type Database struct {
	DB *gorm.DB
}

// Connect establece una conexión a la base de datos PostgreSQL
func (d *Database) Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DSN not found")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %w", err)
	}

	// Configurar el pool de conexiones
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("could not get database instance: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Hacer ping a la base de datos para verificar la conexión
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping the database: %w", err)
	}

	d.DB = db
	return db, nil
}

// NewDatabaseInstance crea una nueva instancia de Database
func NewDatabaseInstance() NewDatabase {
	return &Database{}
}
