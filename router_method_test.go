package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterMethod(t *testing.T) {
	t.Run("method matches", func(t *testing.T) {
		router := New()
		router.Get("/", newHandler())

		req, _ := http.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("method not found", func(t *testing.T) {
		router := New()
		router.Get("/", newHandler())

		req, _ := http.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotImplemented {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusNotImplemented, status)
		}
	})
}
