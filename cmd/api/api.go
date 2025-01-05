package api

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"log"
	"net/http"

	"github.com/matthewjamesboyle/golang-interview-prep/internal/config"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/user"
)

var env = config.Envs
var DATABASE_URL = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", env.DB_TYPE, env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)

func StartServer() {
	// runMigrations()

	h := user.Handler{Svc: *svc}

	http.HandleFunc("/user", h.AddUser)
	http.HandleFunc("/test", h.Test)
	log.Printf("starting http server on http://localhost%s/test\n", env.PORT)
	log.Fatal(http.ListenAndServe(env.PORT, nil))
}

// func runMigrations() {
// 	// Database connection string
// 	fmt.Println(env.MIGRATION_PATH)
// 	dbURL := DATABASE_URL

// 	db, err := sql.Open("postgres", dbURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("db is down: %s", err)
// 	}
// 	fmt.Printf("db.Stats().InUse: %v\n", db.Stats().OpenConnections)
// 	// Create a new instance of the PostgreSQL driver for migrate
// 	driver, err := postgres.WithInstance(db, &postgres.Config{})
// 	if err != nil {
// 		log.Fatalf("failed to create db driver %v", err)
// 	}

// 	m, err := migrate.NewWithDatabaseInstance(env.MIGRATION_PATH, env.DB_NAME, driver)
// 	if err != nil {
// 		log.Fatalf("failed to initialize migrations %v", err)
// 	}

// 	err = m.Up()
// 	if err != nil && err != migrate.ErrNoChange {
// 		log.Fatal(err)
// 	}

// 	log.Println("Database migration complete.")
// }
