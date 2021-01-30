package service

import (
	"io"
	"net/http"

	"github.com/alekssum/todo/internal/handler"

	"github.com/gorilla/mux"
)

func (s *service) registerRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(s.loggingMiddleware)

	r.HandleFunc("/health", handler.HealthCheck)

	r.HandleFunc(
		"/articles/{category}/{id:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "Ok")
		},
	).Name("article")

	return r
}
