package db

import "fmt"

type Config struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func (cfg *Config) ConnectionURI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}
