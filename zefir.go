package zefir

import (
	"log"
	"os"
	"fmt"
	"net/http"
	"strconv"
)

const (
	PORT    int = 8080
	VERSION string = "0.0.1"
)

type Application struct {
	*Router
}

func Create() *Application  {
	log.Printf("zefir v%s", VERSION)

	router:= NewRouter()
	return &Application{&router}
}

func (a *Application) GetRouter() *Router {
	return a.Router
}


func (a *Application) Start() {
	port:= PORT
	if p := os.Getenv("PORT"); p != ""{
		port, _ = strconv.Atoi(p)
	}

	log.Printf("Server started on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), a)
}


func (a *Application) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	//TODO: Add default handlers
	matchedRoute:= a.GetRouter().Get404Route()

	for _, route := range a.GetRouter().GetRoutes() {
		if (route.method == req.Method || "*" == route.method) && route.pattern.MatchString(req.URL.Path) {
			matchedRoute = route
		}
	}
	a.HandleRoute(res, req, matchedRoute)
}

func (a *Application) HandleRoute(res http.ResponseWriter, req *http.Request, route Route) {
	context := NewContext()
	context.SetRoute(&route)
	context.SetRequest(req)
	context.SetResponse(res)

	handler:=context.GetRoute().GetHandler()
	handler(context)
}


