package service

import (
	"fmt"
	"os"

	"github.com/alekssum/todo/internal/db"

	"github.com/alekssum/todo/internal/logging"
)

func NewConfig() *config {
	return &config{
		HTTP: HTTP{Host: ":", Port: "8090"},
		Log: &logging.Config{
			logging.DebugLevelName,
		},
		DB: &db.Config{},
	}
}

type config struct {
	HTTP HTTP
	DB   *db.Config
	Log  *logging.Config
}

type HTTP struct {
	Host string
	Port string
}

func (h *HTTP) Address() string {
	return fmt.Sprintf("%s%s", h.Host, h.Port)
}

const (
	dbHost = "TODO_DATABASE_HOST"
	dbPort = "TODO_DATABASE_PORT"
	dbName = "TODO_DATABASE_NAME"
	dbUser = "TODO_DATABASE_USER"
	dbPass = "TODO_DATABASE_PASS"
)

func (c *config) Init(file string) error {

	if v, ok := os.LookupEnv(dbHost); ok {
		c.DB.Host = v
	}

	if v, ok := os.LookupEnv(dbPort); ok {
		c.DB.Port = v
	}

	if v, ok := os.LookupEnv(dbName); ok {
		c.DB.Name = v
	}

	if v, ok := os.LookupEnv(dbUser); ok {
		c.DB.User = v
	}

	if v, ok := os.LookupEnv(dbPass); ok {
		c.DB.Pass = v
	}

	return nil
}
