package service

import "fmt"

func NewConfig() *config {
	return &config{
		HTTP: struct {
			Host string
			Port string
		}{Host: ":", Port: "8090"},
	}
}

type config struct {
	HTTP struct {
		Host string
		Port string
	}
	DB struct {
		User string
		Pass string
		Host string
		Port string
		Name string
	}
}

func (c *config) Address() string {
	return fmt.Sprintf("%s%s", c.HTTP.Host, c.HTTP.Port)
}

func (c *config) Init(file string) error {

	return nil
}
