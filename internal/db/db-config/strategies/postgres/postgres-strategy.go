package postgres_strategy

import "fmt"

type PostgresStrategy struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func (p PostgresStrategy) DriverName() string {
	return "postgres"
}

func (p PostgresStrategy) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DbName, p.SSLMode,
	)
}
