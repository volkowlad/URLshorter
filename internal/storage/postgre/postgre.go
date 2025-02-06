package postgre

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234qwer"
	dbname   = "postgres"
	sslmode  = "disable"
)

type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitPostgre(cfg ConfigDB) (*sql.DB, error) {
	dsn := fmt.Sprintf(`
		host=%s 
		port=%d
		user=%s 
		password=%s 
		dbname=%s
		ssmode=%S
		`, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Errorf("cannot open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Errorf("cannot ping database: %w", err)
	}

	return db, nil
}
