package mysql_strategy

import "fmt"

type MySQLStrategy struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (m MySQLStrategy) DSN() string {
	return "mysql"
}

func (m MySQLStrategy) DriverName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		m.User, m.Password, m.Host, m.Port, m.DbName,
	)
}
