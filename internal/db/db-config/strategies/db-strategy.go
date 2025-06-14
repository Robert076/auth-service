package db_strategy

type DBStrategy interface {
	DriverName() string
	DSN() string
}
