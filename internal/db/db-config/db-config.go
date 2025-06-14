package db_config

import (
	"database/sql"
	"fmt"
	"os"

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
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}, nil
}

func (cfg *DBConfig) InitDB() (*sql.DB, error) {
	var driver string
	var dsn string

	switch cfg.Type {
	case Postgres:
		driver = "postgres"
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
		)
	case MySQL:
		driver = "mysql"
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
		)
	default:
		return nil, fmt.Errorf("unsupported db type: %s", cfg.Type)
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping failed: %v", err)
	}

	return db, nil
}
