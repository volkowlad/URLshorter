package postgre

import (
	"database/sql"
	"errors"
	"fmt"
	"url_rest_api/internal/storage"
)

type Storage struct {
	db *sql.DB
}
type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitPostgre(cfg ConfigDB) (*Storage, error) {
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

	return &Storage{db: db}, nil
}

// возвращаем индекс созданной записи
func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("cannot SaveURL: %w", err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil {
		// TODO: refactor for PostgreSQL constraints (уникальность alias)
		return 0, fmt.Errorf("cannot SaveURL: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("cannot SaveURL: %w", err)
	}

	return id, nil
}

func (s *Storage) GetURL(aliasGet string) (string, error) {
	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias=?")
	if err != nil {
		return "", fmt.Errorf("cannot GetURL: %w", err)
	}

	var resURL string

	err = stmt.QueryRow(aliasGet).Scan(&resURL)
	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrURLNotFound
	}
	if err != nil {
		return "", fmt.Errorf("cannot GetURL: %w", err)
	}

	return resURL, nil
}

func (s *Storage) DeleteURL(aliasDelete string) error {
	stmt, err := s.db.Prepare("DELETE FROM url WHERE alias=?")
	if err != nil {
		return fmt.Errorf("cannot DeleteURL: %w", err)
	}

	_, err = stmt.Exec(aliasDelete)
	if err != nil {
		return fmt.Errorf("cannot DeleteURL: %w", err)
	}

	return nil
}
