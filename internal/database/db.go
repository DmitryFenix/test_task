package database

import (
	"fmt"
	"log"
	"time"
	"qa-api/internal/config"
	"qa-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init initializes the database connection with retry logic
func Init(cfg *config.Config) error {
	var err error
	maxRetries := 10
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			// Auto-migrate models (as a fallback, but we use goose migrations)
			if err := DB.AutoMigrate(&models.Question{}, &models.Answer{}); err != nil {
				log.Printf("Warning: auto-migrate failed: %v", err)
			}

			log.Println("Database connection established")
			return nil
		}

		if i < maxRetries-1 {
			log.Printf("Failed to connect to database (attempt %d/%d): %v. Retrying in %v...", i+1, maxRetries, err, retryDelay)
			time.Sleep(retryDelay)
		}
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}





