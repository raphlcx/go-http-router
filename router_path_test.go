package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterPath(t *testing.T) {
	t.Run("match on same path but different method", func(t *testing.T) {
		router := New()
		router.Get("/", newHandler())
		router.Post("/", newHandler())

		req, _ := http.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}

		req, _ = http.NewRequest("GET", "/", nil)
		rr = httptest.NewRecorder()

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}

		req, _ = http.NewRequest("PUT", "/", nil)
		rr = httptest.NewRecorder()

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotImplemented {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusNotImplemented, status)
		}
	})

	t.Run("match on method first and then path", func(t *testing.T) {
		router := New()
		router.Get("/api/users", newHandler())
		router.Post("/api/products", newHandler())
		router.Options("/api/", newHandler())

		req, _ := http.NewRequest("OPTIONS", "/api/products", nil)
		rr := httptest.NewRecorder()

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("no path matches", func(t *testing.T) {
		router := New()
		router.Get("/api/users", newHandler())

		req, _ := http.NewRequest("GET", "/api/products", nil)
		rr := httptest.NewRecorder()

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusNotFound, status)
		}
	})
}
