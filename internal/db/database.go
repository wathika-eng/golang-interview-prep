package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/config"
)

var (
	env    = config.Envs
	DbPool *pgxpool.Pool
	ctx    = context.Background()
)

// StartDB initializes the database connection pool and prepares the queries.
func StartDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_NAME)
	fmt.Println(connStr)
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("failed to parse database configuration: %v", err)
	}

	// Configure connection pooling settings
	poolConfig.MaxConns = 10                       // Maximum number of connections in the pool
	poolConfig.MinConns = 2                        // Minimum number of connections in the pool
	poolConfig.MaxConnLifetime = time.Hour         // Maximum lifetime of a connection
	poolConfig.MaxConnIdleTime = 30 * time.Minute  // Maximum idle time for a connection
	poolConfig.HealthCheckPeriod = 5 * time.Minute // Health check interval

	DbPool, err = pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Test the connection
	if err = DbPool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping the database: %v", err)
	}
	createTable(DbPool)
	log.Println("Database connection pool initialized successfully.")
}

// CloseDBConnection closes the database connection pool.
func CloseDBConnection() {
	if DbPool != nil {
		DbPool.Close()
		log.Println("Database connection pool closed.")
	}
}

// creates users table upon first initialization of the DB
//
// will later use migrate package from Golang
func createTable(db *pgxpool.Pool) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		ID UUID NOT NULL,
		work_id SERIAL PRIMARY KEY NOT NULL,
		username VARCHAR(20) NOT NULL UNIQUE,
		email VARCHAR(255) NOT NULL UNIQUE,
		phone_number VARCHAR(20) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	
	-- Create indexes on username and email
	CREATE INDEX IF NOT EXISTS idx_username ON users (username);
	CREATE INDEX IF NOT EXISTS idx_email ON users (email);
	`
	result, err := db.Exec(ctx, query)
	if err != nil {
		log.Fatalf("table creation failed: %v", err)
	}
	log.Print(result)
}
