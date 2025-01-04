package main

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"log"
	"net/http"

	"github.com/matthewjamesboyle/golang-interview-prep/internal/config"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/user"
)

var env = config.Envs
var DATABASE_URL = fmt.Sprintf("%s://%s:%s@%s:%s/%s", env.DB_TYPE, env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)

func main() {
	runMigrations()
	svc, err := user.NewService(env.DB_USER, env.DB_PASSWORD)
	if err != nil {
		log.Fatal(err)
	}

	h := user.Handler{Svc: *svc}

	http.HandleFunc("/user", h.AddUser)
	http.HandleFunc("/test", h.Test)
	log.Printf("starting http server on http://localhost%s\n", env.PORT)
	log.Fatal(http.ListenAndServe(env.PORT, nil))
}

func runMigrations() {
	// Database connection string
	// fmt.Println(DATABASE_URL)
	dbURL := DATABASE_URL

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("db is down: %s", err)
	}
	fmt.Printf("db.Stats().InUse: %v\n", db.Stats().InUse)
	// Create a new instance of the PostgreSQL driver for migrate
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Database migration complete.")
}
