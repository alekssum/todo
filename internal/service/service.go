package service

import (
	"log"
	"net/http"

	"github.com/alekssum/todo/internal/logging"

	_ "github.com/lib/pq"
)

func New(cfg *config) *service {
	return &service{
		cfg: cfg,
	}
}

type service struct {
	cfg *config
	l   logging.Logger
}

func (s *service) Start() error {
	s.configureLogger()

	//db, err := db.New(s.cfg.DB)
	//if err != nil {
	//	return errors.Wrap(err, "service starting")
	//}
	//defer db.Close()

	r := s.registerRoutes()
	httpAddr := s.cfg.HTTP.Address()
	log.Println("Listen and serve:", httpAddr)
	return http.ListenAndServe(httpAddr, r)
}

func (s *service) configureLogger() {
	l := logging.New(s.cfg.Log)
	s.l = l
	log.Println("Log level:", s.cfg.Log.LogLevel)
}
