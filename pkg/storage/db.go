package storage

import (
	"database/sql"
	"fmt"
	"invest-tracker/pkg/config"

	_ "github.com/lib/pq"
)

type Database interface {
	Init() error
	Close() error
	Instance() *sql.DB
}

type PostgresDB struct {
	db *sql.DB
}

func (pg *PostgresDB) Init() error {
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	pg.db, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	return pg.db.Ping()
}

func (pg *PostgresDB) Close() error {
	if pg.db != nil {
		return pg.db.Close()
	}
	return nil
}

func (pg *PostgresDB) Instance() *sql.DB {
	return pg.db
}
