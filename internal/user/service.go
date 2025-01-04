package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/config"
)

var env = config.Envs
var DATABASE_URL = fmt.Sprintf("%s://%s:%s@%s:%s/%s", env.DB_TYPE, env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)

type service struct {
	dbUser     string
	dbPassword string
}

func NewService(dbUser, dbPassword string) (*service, error) {
	if dbUser == "" || dbPassword == "" {
		return nil, errors.New("empty field found")
	}
	return &service{dbUser: dbUser, dbPassword: dbPassword}, nil
}

type User struct {
	Name     string
	Password string
}

func (s *service) AddUser(u User) (string, error) {
	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalf("error connecting to the db: %s\n", err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("db is down: %s", err)
	}
	var id string
	q := "INSERT INTO users (username, password) VALUES ('" + u.Name + "', '" + u.Password + "') RETURNING id"

	err = db.QueryRow(q).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert: %w", err)
	}

	return id, nil
}
