package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	DatabaseURL string
	Port        string
}

// Load reads configuration from environment variables
func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/qa_db?sslmode=disable"),
		Port:        getEnv("PORT", "8080"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}





