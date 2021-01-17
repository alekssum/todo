package service

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func New(cfg *config) *service {
	return &service{cfg}
}

type service struct {
	cfg *config
}

func (s *service) Start() error {
	//db, err := s.connectDB()
	//if err != nil {
	//	return err
	//}
	//defer db.Close()

	s.registerRoutes()

	return http.ListenAndServe(s.cfg.Address(), nil)
}

func (s *service) connectDB() (*sql.DB, error) {
	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.cfg.DB.Host,
		s.cfg.DB.Port,
		s.cfg.DB.User,
		s.cfg.DB.Pass,
		s.cfg.DB.Name,
	)
	db, err := sql.Open("postgres", connURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
