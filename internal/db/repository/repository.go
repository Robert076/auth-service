package repository

import (
	"database/sql"
	"fmt"
	"time"

	hashing_service "github.com/Robert076/auth-service/internal/service/hashing-service"
	user "github.com/Robert076/auth-service/internal/user"
	_ "github.com/lib/pq"
)

func RegisterUser(db *sql.DB, user user.User) error {
	hashedPassword, err := hashing_service.HashPassword(user.Password)

	if err != nil {
		return fmt.Errorf("could not hash password: %v", err)
	}

	query := `INSERT INTO "Users"("Username", "Email", "Password", "CreatedAt") VALUES($1, $2, $3, $4)`

	if _, err := db.Exec(query, user.Username, user.Email, hashedPassword, time.Now()); err != nil {
		return fmt.Errorf("error inserting new user: %v", err)
	}

	return nil
}
