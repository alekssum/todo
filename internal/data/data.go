package data

import (
	"github.com/alekssum/todo/internal/db"
)

func New(db *db.DB) *data {
	return &data{db}
}

type data struct {
	*db.DB
}

func Init() error {
	return nil
}
