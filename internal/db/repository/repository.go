package repository

import (
	"database/sql"
	"fmt"

	db_config "github.com/Robert076/auth-service/internal/db/db-config"
	_ "github.com/lib/pq"
)

func RegisterUser() error {
	dbConfig := db_config.LoadDBConfig()

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}
}
