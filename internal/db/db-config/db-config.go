package db_config

import (
	"database/sql"
	"fmt"
	"os"

	db_strategy "github.com/Robert076/auth-service/internal/db/db-config/strategies"
	postgres "github.com/Robert076/auth-service/internal/db/db-config/strategies/postgres"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type DbType string

const (
	Postgres DbType = "postgres"
	MySQL    DbType = "mysql"
)

type DBConfig struct {
	Type     DbType
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func LoadDBConfig() (DBConfig, error) {
	var dbType DbType = DbType(os.Getenv("DB_TYPE"))

	if dbType != Postgres && dbType != MySQL {
		return DBConfig{}, fmt.Errorf("invalid db type taken from env: %s", dbType)
	}

	return DBConfig{
		Type:     dbType,
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}, nil
}

func (cfg DBConfig) Strategy() (db_strategy.DBStrategy, error) {
	switch cfg.Type {
	case Postgres:
		return postgres.PostgresStrategy{
			Host: cfg.Host, Port: cfg.Port, User: cfg.User,
			Password: cfg.Password, DbName: cfg.DBName, SSLMode: cfg.SSLMode,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported db type: %s", cfg.Type)
	}
}

func InitDB(strategy db_strategy.DBStrategy) (*sql.DB, error) {
	db, err := sql.Open(strategy.DriverName(), strategy.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping failed: %v", err)
	}

	return db, nil
}
