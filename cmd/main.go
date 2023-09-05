package main

import (
	"log"
	"net/http"

	"dev.raphlcx.com/go/router"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Success"))
	return
}

func main() {
	r := router.New()

	r.Get("/", newHandler())
	// r.Head("/", newHandler())
	// r.Post("/", newHandler())
	// r.Put("/", newHandler())
	// r.Patch("/", newHandler())
	// r.Delete("/", newHandler())
	// r.Connect("/", newHandler())
	// r.Options("/", newHandler())
	// r.Trace("/", newHandler())

	log.Fatal(http.ListenAndServe("localhost:8888", r.Mux()))
}
