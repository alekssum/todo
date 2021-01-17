package data

import "database/sql"

func New(db *sql.DB) *data {
	return &data{db}
}

type data struct {
	*sql.DB
}

func Init() error {
	return nil
}
