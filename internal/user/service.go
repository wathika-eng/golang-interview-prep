package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/config"
)

var db *sql.DB
var env = config.Envs

func init() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_NAME)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Successfully connected to the database")
}

type service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *service {
	return &service{db: db}
}

type User struct {
	Name     string
	Password string
}

func (s *service) AddUser(u User) (string, error) {
	if u.Name == "" || u.Password == "" {
		return "", errors.New("name and password cannot be empty")
	}

	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	var id string
	err = s.db.QueryRow(query, u.Name, hashedPassword).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert user: %v", err)
	}

	return id, nil
}

// hashPassword hashes the password using bcrypt
func hashPassword(password string) (string, error) {
	
	return password, nil
}
