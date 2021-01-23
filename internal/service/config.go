package service

import (
	"fmt"

	"github.com/alekssum/todo/internal/logging"
)

func NewConfig() *config {
	return &config{
		HTTP: HTTP{Host: ":", Port: "8090"},
	}
}

type config struct {
	HTTP HTTP
	DB   DB
	Log  logging.Config
}

type HTTP struct {
	Host string
	Port string
}

func (h *HTTP) Address() string {
	return fmt.Sprintf("%s%s", h.Host, h.Port)
}

type DB struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func (db *DB) ConnectionURI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db.Host,
		db.Port,
		db.User,
		db.Pass,
		db.Name,
	)
}

func (c *config) Init(file string) error {

	return nil
}
