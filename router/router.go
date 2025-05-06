package router

import (
	"context"
	"net/http"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) HandleFunc(method, pattern string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{Method: method, Pattern: pattern, HandlerFunc: handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {

		if req.Method != route.Method {
			continue

		}

		params, ok := match(route.Pattern, req.URL.Path)
		if !ok {
			continue
		}

		ctx := context.WithValue(req.Context(), "params", params)
		route.HandlerFunc(w, req.WithContext(ctx))
		return

	}
	http.NotFound(w, req)
}

func Param(r *http.Request, key string) string {
	params, ok := r.Context().Value("params").(map[string]string)
	if !ok {
		return ""
	}
	return params[key]
}
