package service

import (
	"io"
	"net/http"
)

func (s *service) registerRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Ok")
	})

}
