package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterChainable(t *testing.T) {
	t.Run("smoke", func(t *testing.T) {
		mux := New().
			Get("/abc", newHandler()).
			Post("/def", newHandler()).
			Mux()

		req, _ := http.NewRequest("GET", "/abc", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}

		req, _ = http.NewRequest("POST", "/def", nil)
		rr = httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}

		req, _ = http.NewRequest("GET", "/", nil)
		rr = httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusNotFound, status)
		}

		req, _ = http.NewRequest("PUT", "/", nil)
		rr = httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotImplemented {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusNotImplemented, status)
		}
	})
}
