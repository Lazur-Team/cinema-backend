package pgsql

import (
	"cinema/internal/config"
	"cinema/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Storage contains an instance of *gorm.DB to interact with the database.
type Storage struct {
	DB *gorm.DB
}

// New initializes a new database connection using GORM.
// It takes a Config structure containing the connection parameters.
func New(cfg *config.Config) *Storage {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	// Configuration GORM
	gormConfig := &gorm.Config{}

	// Connection to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		panic("Failed to connect to database using GORM:" + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get generic database object: %w" + err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		panic("Failed to ping database: %w" + err.Error())
	}

	db.AutoMigrate(&models.Todo{})

	return &Storage{
		DB: db,
	}
}

// Close closes the database connection.
// This method should be called at the end of the program.
func (s *Storage) Close() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return fmt.Errorf("Failed to get generic database object: %w", err)
	}
	return sqlDB.Close()
}
