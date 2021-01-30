package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

func New(cfg *Config) (*DB, error) {
	connURI := cfg.ConnectionURI()
	db, err := sql.Open("postgres", connURI)
	if err != nil {
		return nil, errors.Wrap(err, "connectDB")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "connectDB")
	}

	return &DB{db}, nil
}

type DB struct {
	*sql.DB
}

func (db *DB) Close() error {
	return db.DB.Close()
}
