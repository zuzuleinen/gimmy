package infrastructure

import "net/http"

type Router struct {
	routes map[string]Route
}

func NewRouter() *Router {
	var r Router
	r.routes = make(map[string]Route)
	return &r
}

type Route struct {
	Method     string
	Path       string
	HandleFunc func(w http.ResponseWriter, r *http.Request) error
}

func (r *Router) AddRoute(route Route) {
	r.routes[route.Method+route.Path] = route
}

func (r *Router) FindRoute(method, path string) (Route, bool) {
	rt, ok := r.routes[method+path]
	return rt, ok
}
