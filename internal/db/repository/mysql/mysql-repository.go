package mysql_repository

import (
	"database/sql"
	"fmt"
	"time"

	hashing_service "github.com/Robert076/auth-service/internal/service/hashing-service"
	user "github.com/Robert076/auth-service/internal/user"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) RegisterUser(user user.User) error {
	hashedPassword, err := hashing_service.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("could not hash password: %v", err)
	}

	query := `INSERT INTO Users (Username, Email, Password, CreatedAt) VALUES (?, ?, ?, ?)`
	_, err = r.db.Exec(query, user.Username, user.Email, hashedPassword, time.Now())
	if err != nil {
		return fmt.Errorf("mysql insert error: %v", err)
	}
	return nil
}
