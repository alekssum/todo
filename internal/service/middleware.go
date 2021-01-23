package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *service) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("METHOD[%s] URI: %s", r.Method, r.RequestURI)
		switch r.Method {
		case http.MethodPost:
			b, _ := ioutil.ReadAll(r.Body)
			fmt.Print("BODY:")
			fmt.Println(string(b))
		}

		next.ServeHTTP(w, r)
	})
}
