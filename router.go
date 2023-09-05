package router

import (
	"net/http"
)

type router struct {
	mux *http.ServeMux
}

func (r *router) Get(path string, handler http.Handler) {
	r.route(http.MethodGet, path, handler)
}

func (r *router) Head(path string, handler http.Handler) {
	r.route(http.MethodHead, path, handler)
}

func (r *router) Post(path string, handler http.Handler) {
	r.route(http.MethodPost, path, handler)
}

func (r *router) Put(path string, handler http.Handler) {
	r.route(http.MethodPut, path, handler)
}

func (r *router) Patch(path string, handler http.Handler) {
	r.route(http.MethodPatch, path, handler)
}

func (r *router) Delete(path string, handler http.Handler) {
	r.route(http.MethodDelete, path, handler)
}

func (r *router) Connect(path string, handler http.Handler) {
	r.route(http.MethodConnect, path, handler)
}

func (r *router) Options(path string, handler http.Handler) {
	r.route(http.MethodOptions, path, handler)
}

func (r *router) Trace(path string, handler http.Handler) {
	r.route(http.MethodTrace, path, handler)
}

func (r *router) route(method string, path string, handler http.Handler) {
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			w.Header().Set(http.CanonicalHeaderKey("allow"), method)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, req)
	})
}

func (r *router) Mux() *http.ServeMux {
	return r.mux
}

func New() *router {
	return &router{mux: http.NewServeMux()}
}
