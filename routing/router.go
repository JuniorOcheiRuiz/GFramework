package routing

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Router struct {
	Server *http.Server
	Routes []*Route
}

type RouterOptions struct {
	Port int
	SSL  bool
}

func NewRouter(options *RouterOptions) *Router {
	server := &http.Server{}
	router := NewRouterFromServer(server)

	// server configuration
	server.Addr = ":" + strconv.Itoa(options.Port)
	server.Handler = NewRequestHandler(router)

	return router
}

func NewRouterFromServer(server *http.Server) *Router {
	router := &Router{}
	router.Server = server
	return router
}

func (r *Router) FindRouteByPath(path string) (*Route, RouteParams) {

	for _, route := range r.Routes {
		if isFound, params := route.Match(path); isFound {
			return route, params
		}
	}

	return nil, nil
}

func (r *Router) FindRouteByName(name string) *Route {
	panic("FindRouteByName: Not implemented yet.")
}

func (r *Router) Run() {

	fmt.Println("Servidor escuchando en el puerto " + r.Server.Addr)
	log.Fatal(r.Server.ListenAndServe())
}
