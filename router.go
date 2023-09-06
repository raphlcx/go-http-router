package router

import (
	"net/http"
)

type router struct {
	muxes map[string]*http.ServeMux
}

func (r *router) Get(path string, handler http.Handler) *router {
	return r.route(http.MethodGet, path, handler)
}

func (r *router) Head(path string, handler http.Handler) *router {
	return r.route(http.MethodHead, path, handler)
}

func (r *router) Post(path string, handler http.Handler) *router {
	return r.route(http.MethodPost, path, handler)
}

func (r *router) Put(path string, handler http.Handler) *router {
	return r.route(http.MethodPut, path, handler)
}

func (r *router) Patch(path string, handler http.Handler) *router {
	return r.route(http.MethodPatch, path, handler)
}

func (r *router) Delete(path string, handler http.Handler) *router {
	return r.route(http.MethodDelete, path, handler)
}

func (r *router) Connect(path string, handler http.Handler) *router {
	return r.route(http.MethodConnect, path, handler)
}

func (r *router) Options(path string, handler http.Handler) *router {
	return r.route(http.MethodOptions, path, handler)
}

func (r *router) Trace(path string, handler http.Handler) *router {
	return r.route(http.MethodTrace, path, handler)
}

func (r *router) route(method string, path string, handler http.Handler) *router {
	mux, ok := r.muxes[method]
	if !ok {
		mux = http.NewServeMux()
		r.muxes[method] = mux
	}
	mux.Handle(path, handler)
	return r
}

func (r *router) Mux() *http.ServeMux {
	routerMux := http.NewServeMux()

	routerMux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		mux, ok := r.muxes[req.Method]
		if !ok {
			http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
			return
		}
		mux.ServeHTTP(w, req)
	})

	return routerMux
}

func New() *router {
	return &router{muxes: make(map[string]*http.ServeMux)}
}
