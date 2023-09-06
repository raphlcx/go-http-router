package router

import (
	"net/http"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	return
}
