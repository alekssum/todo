package main

import (
	"log"

	"github.com/alekssum/todo/internal/service"
)

func main() {
	cfg := service.NewConfig()
	if err := cfg.Init(""); err != nil {
		log.Fatal(err)
	}

	s := service.New(cfg)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
