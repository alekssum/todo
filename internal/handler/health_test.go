package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(HealthCheck)
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HealthCheck returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `Ok`
	if rr.Body.String() != expected {
		t.Errorf("HealthCheck returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
