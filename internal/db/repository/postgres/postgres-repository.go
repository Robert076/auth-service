package postgres_repository

import (
	"database/sql"
	"fmt"
	"time"

	hashing_service "github.com/Robert076/auth-service/internal/service/hashing-service"
	user "github.com/Robert076/auth-service/internal/user"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	Db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{Db: db}
}

func (r *PostgresRepository) RegisterUser(user user.RegisterUserDTO) error {
	hashedPassword, err := hashing_service.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("could not hash password: %v", err)
	}

	insertQuery := `INSERT INTO "Users"("Username", "Email", "Password", "CreatedAt") VALUES($1, $2, $3, $4)`
	_, err = r.Db.Exec(insertQuery, user.Username, user.Email, hashedPassword, time.Now())
	if err != nil {
		return fmt.Errorf("postgres insert error: %v", err)
	}
	return nil
}

func (r *PostgresRepository) LoginUser(user user.LoginUserDTO) error {
	getPasswordQuery := `SELECT "Password" FROM "Users" WHERE "Email" = $1`

	var storedPassword string
	err := r.Db.QueryRow(getPasswordQuery, user.Email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no user found with email: %s", user.Email)
		}
		return fmt.Errorf("error getting user's password from db: %v", err)
	}

	if !hashing_service.CompareHash(storedPassword, user.Password) {
		return fmt.Errorf("password doesn't match")
	}
	return nil
}

func (r *PostgresRepository) SetSessionToken(email string, sessionToken string) error {
	setSessionTokenQuery := `UPDATE "Users" SET "SessionToken" = $1 WHERE "Email" = $2`

	_, err := r.Db.Exec(setSessionTokenQuery, sessionToken, email)
	if err != nil {
		return fmt.Errorf("error updating session token for user %s, got error: '%v'", email, err)
	}
	return nil
}

func (r *PostgresRepository) SetCsrfToken(email string, csrfToken string) error {
	setCsrfTokenQuery := `UPDATE "Users" SET "CsrfToken" = $1 WHERE "Email" = $2`

	_, err := r.Db.Exec(setCsrfTokenQuery, csrfToken, email)
	if err != nil {
		return fmt.Errorf("error updating session token for user %s, got error: '%v'", email, err)
	}
	return nil
}
