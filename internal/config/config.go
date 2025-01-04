package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds configuration values for the application
type Config struct {
	DB_TYPE     string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	PORT        string
}

var Envs = initConfig()

// InitConfig initializes and returns the application configuration
func initConfig() *Config {
	// Load .env file, log a fatal error if it fails
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("Warning: No .env file found or failed to load. Using default environment variables.")
	}

	return &Config{
		DB_TYPE:     getEnv("DB_TYPE", "postgres"),
		DB_USER:     getEnv("DB_USER", "postgres"),
		DB_PASSWORD: getEnv("DB_PASSWORD", ""),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_NAME:     getEnv("DB_NAME", "interview"),
		PORT:        getEnv("PORT", ":8080"),
	}
}

// getEnv retrieves an environment variable or returns the fallback string
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
