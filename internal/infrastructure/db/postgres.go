package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
	"rockstar-corp/internal/config"
)

type DB struct {
	Conn *sql.DB
}

func NewPostgres(cfg config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Ping to ensure it's working
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}

func (d *DB) Close() {
	_ = d.Conn.Close()
}
