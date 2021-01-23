package service

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/alekssum/todo/internal/alogsubscriber"

	"github.com/alekssum/todo/internal/logging"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
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

	//db, err := s.connectDB()
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
	logSubscribers := make([]logging.LogSubscriber, 0)
	logSubscribers = append(logSubscribers, &alogsubscriber.MQ{})
	l := logging.New(&s.cfg.Log)
	l.Subscribe(logSubscribers)
	s.l = l
}

func (s *service) connectDB() (*sql.DB, error) {
	connURI := s.cfg.DB.ConnectionURI()
	db, err := sql.Open("postgres", connURI)
	if err != nil {
		return nil, errors.Wrap(err, "connectDB")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "connectDB")
	}

	return db, nil
}
