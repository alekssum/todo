package service

import (
	"fmt"

	"github.com/alekssum/todo/internal/db"

	"github.com/alekssum/todo/internal/logging"
)

func NewConfig() *config {
	return &config{
		HTTP: HTTP{Host: ":", Port: "8090"},
		Log: &logging.Config{
			logging.DebugLevelName,
		},
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

func (c *config) Init(file string) error {

	return nil
}
