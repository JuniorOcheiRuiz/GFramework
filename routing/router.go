package routing

import (
	"log"
	"net/http"
	"strconv"
)

type Router struct {
	Server *http.Server
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

func (r *Router) FindRouteByPath(path string) (*Route, map[string]interface{}) {

	return nil, nil
}

func (r *Router) FindRouteByName(name string) *Route {
	return nil
}

func (r *Router) Run() {

	log.Fatal(r.Server.ListenAndServe())
}
