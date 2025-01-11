package db

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testDBPool *pgxpool.Pool

// SetupTestDB initializes a test database connection.
func SetupTestDB() *pgxpool.Pool {
	connStr := "host=localhost port=5432 user=testuser password=testpassword dbname=testdb sslmode=disable"

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("failed to parse test database configuration: %v", err)
	}

	// Configure test-specific connection pooling settings
	poolConfig.MaxConns = 5
	poolConfig.MinConns = 1
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 15 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	testDBPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("failed to connect to the test database: %v", err)
	}

	// Test the connection
	if err = testDBPool.Ping(context.Background()); err != nil {
		log.Fatalf("failed to ping the test database: %v", err)
	}

	return testDBPool
}

// TestStartDB tests the StartDB function.
func TestStartDB(t *testing.T) {
	// Set up test database
	testDBPool = SetupTestDB()
	defer testDBPool.Close()

	// Call createTable directly to simulate table creation in the test database
	createTable(testDBPool)

	// Verify table creation
	query := `
	SELECT EXISTS (
		SELECT FROM information_schema.tables
		WHERE table_name = 'users'
	);`
	var tableExists bool
	err := testDBPool.QueryRow(context.Background(), query).Scan(&tableExists)
	if err != nil {
		t.Fatalf("failed to verify table existence: %v", err)
	}

	if !tableExists {
		t.Fatalf("users table was not created")
	}

	t.Log("StartDB test passed: users table exists")
}

// TestCloseDBConnection tests the CloseDBConnection function.
func TestCloseDBConnection(t *testing.T) {
	// Ensure the connection pool is set
	testDBPool = SetupTestDB()
	CloseDBConnection()

	// Attempt to use a closed connection pool
	err := testDBPool.Ping(context.Background())
	if err == nil {
		t.Fatalf("expected error when pinging a closed connection pool, but got none")
	}

	t.Log("CloseDBConnection test passed: connection pool closed successfully")
}
