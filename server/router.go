package server

import (
	"strings"
)

type Router struct {
	routes map[string]Route
}

func NewRouter() Router {
	return Router{routes: make(map[string]Route)}
}

func (r *Router) Add(path string, method string, f func(res Response, req Request) *Response) {
	pattern := strings.Split(path, ":")[0]
	pattern = strings.TrimSuffix(pattern, "/")
	if r.routes[pattern].Pattern == "" {
		r.routes[pattern] = NewRoute(path, method, f)
	} else {
		r.routes[pattern].Add(path, method, f)
	}

}

func (r *Router) Find(path, method string) (*Subset, func(res Response, req Request) *Response) {
	pattern := strings.Split(path, "/")[1]
	route := r.routes["/"+pattern]
	if route.Subset == nil {
		return nil, nil
	}
	return route.Search(path, method)
}
