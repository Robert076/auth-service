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

	query := `INSERT INTO "Users"("Username", "Email", "Password", "CreatedAt") VALUES($1, $2, $3, $4)`
	_, err = r.Db.Exec(query, user.Username, user.Email, hashedPassword, time.Now())
	if err != nil {
		return fmt.Errorf("postgres insert error: %v", err)
	}
	return nil
}
