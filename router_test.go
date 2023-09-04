package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}

func TestRouter(t *testing.T) {
	t.Run("GET ok", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Get("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("GET not ok", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Get("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusMethodNotAllowed, status)
		}
	})

	t.Run("POST ok", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Post("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("POST not ok", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Post("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusMethodNotAllowed, status)
		}
	})

	t.Run("PUT ok", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Put("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("PUT not ok", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Get("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusMethodNotAllowed, status)
		}
	})

	t.Run("PATCH ok", func(t *testing.T) {
		req, _ := http.NewRequest("PATCH", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Patch("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("PATCH not ok", func(t *testing.T) {
		req, _ := http.NewRequest("PATCH", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Put("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusMethodNotAllowed, status)
		}
	})

	t.Run("DELETE ok", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Delete("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("DELETE not ok", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Post("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusMethodNotAllowed, status)
		}
	})

	t.Run("OPTIONS ok", func(t *testing.T) {
		req, _ := http.NewRequest("OPTIONS", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Options("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusOK, status)
		}
	})

	t.Run("OPTIONS not ok", func(t *testing.T) {
		req, _ := http.NewRequest("OPTIONS", "/", nil)
		rr := httptest.NewRecorder()

		router := New()
		router.Patch("/", newHandler())

		router.Mux().ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Expecting status code %d, got %d instead", http.StatusMethodNotAllowed, status)
		}
	})
}
