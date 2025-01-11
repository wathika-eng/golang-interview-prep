package models

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/db"
)

var ctx = context.Background()

type User struct {
	ID          uuid.UUID `json:"id"`
	WorkID      int       `json:"work_id" binding:"required"`
	UserName    string    `json:"username" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	PhoneNumber string    `json:"phone_number" binding:"required,e164"`
	Password    string    `json:"password" binding:"required,min=8"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateUser(u *User) (int, error) {
	u.ID = uuid.Must(uuid.NewV7())
	u.CreatedAt = time.Now()
	query := `
		INSERT INTO users (id, work_id, username, email, phone_number, password, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := db.DbPool.Exec(ctx, query, u.ID, u.WorkID, u.UserName, u.Email, u.PhoneNumber, u.Password, u.CreatedAt)
	if err != nil {
		return 0, errors.New("error inserting user: " + err.Error())
	}
	return u.WorkID, nil
}

func GetUser(workID int) (*User, error) {
	var user User
	query := `
	SELECT  id, work_id, username, email, phone_number, created_at, updated_at 
	FROM users WHERE work_id = $1
	`
	results := db.DbPool.QueryRow(ctx, query, workID)
	err := results.Scan(
		&user.ID,
		&user.WorkID,
		&user.UserName,
		&user.Email,
		&user.PhoneNumber,
		// &user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("error scanning user: " + err.Error())
	}
	return &user, nil
}

func GetUsers() ([]User, error) {
	users := make([]User, 0, 10)
	query := `
	SELECT id, work_id, username, email, phone_number, created_at, updated_at FROM users
	`
	results, err := db.DbPool.Query(ctx, query)
	if err != nil {
		return nil, errors.New("error fetching users")
	}
	defer results.Close()
	for results.Next() {
		var user User
		err := results.Scan(
			&user.ID,
			&user.WorkID,
			&user.UserName,
			&user.Email,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New("error scanning user: " + err.Error())
		}
		users = append(users, user)
	}
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("users not found")
		}
		return nil, errors.New("error scanning user: " + err.Error())
	}

	return users, nil
}
