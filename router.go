package zefir

import (
	"regexp"
)

type Router struct {
	routes []Route
	route404 Route
}


// Creates a new router component instance
func NewRouter() Router {
	r:=Router{routes: make([]Route, 0)}

	r.route404 = Route{
		"*",
		"404",
		regexp.MustCompile("/$"),
		func(context Context) {
			context.GetResponse().SetStatusCode(404)
			context.GetResponse().SetText("Not found")
		},
	}

	return r
}


// Returns all routes available in router
func (r *Router) GetRoutes() []Route {
	return r.routes
}

func (r *Router) Set404Route(route Route){
	r.route404 = route
}

func (r *Router) Get404Route() Route {
	return r.route404
}

func (r *Router) Add(method string, name string, pattern string, handler Handler) {
	route := Route{method, name, regexp.MustCompile(pattern), handler}
	r.routes = append(r.routes, route)
}

// Adds different HTTP methods route
func (r *Router) Get(name string, pattern string, handler Handler) {
	r.Add("GET", name, pattern, handler)
}

func (r *Router) Patch(name string, pattern string, handler Handler) {
	r.Add("PATCH", name, pattern, handler)
}

func (r *Router) Post(name string, pattern string, handler Handler) {
	r.Add("POST", name, pattern, handler)
}

func (r *Router) Put(name string, pattern string, handler Handler) {
	r.Add("PUT", name, pattern, handler)
}

func (r *Router) Delete(name string, pattern string, handler Handler) {
	r.Add("DELETE", name, pattern, handler)
}

func (r *Router) Options(name string, pattern string, handler Handler) {
	r.Add("OPTIONS", name, pattern, handler)
}

func (r *Router) Head(name string, pattern string, handler Handler) {
	r.Add("HEAD", name, pattern, handler)
}

func (r *Router) All(name string, pattern string, handler Handler) {
	r.Add("*", name, pattern, handler)
}
