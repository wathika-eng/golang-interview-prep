// package Config holds configuration values for the application
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds configuration values for the application
type Config struct {
	POSTGRES_TYPE     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_NAME     string
	PORT              string
	REDIS_URL         string
	MIGRATION_PATH    string
}

// get environment variables from .env file
var Envs = initConfig()

// InitConfig initializes and returns the application configuration
func initConfig() *Config {
	// Load .env file, log a fatal error if it fails
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found or failed to load. Using default environment variables.")
	}

	return &Config{
		POSTGRES_TYPE:     getEnv("POSTGRES_TYPE", "postgres"),
		POSTGRES_USER:     getEnv("POSTGRES_USER", "postgres"),
		POSTGRES_PASSWORD: getEnv("POSTGRES_PASSWORD", ""),
		POSTGRES_HOST:     getEnv("POSTGRES_HOST", "postgres"),
		POSTGRES_PORT:     getEnv("POSTGRES_PORT", "5432"),
		POSTGRES_NAME:     getEnv("POSTGRES_NAME", "postgres"),
		PORT:              getEnv("PORT", ":8080"),
		REDIS_URL:         getEnv("REDIS_URL", "localhost:6379"),
		MIGRATION_PATH:    getEnv("MIGRATION_PATH", "file://internal/migrations"),
	}
}

// getEnv retrieves an environment variable or returns the fallback string
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
