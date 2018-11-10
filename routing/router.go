package routing

import (
	"log"
	"net/http"
	"strconv"

	_http "github.com/juniorocheiruiz/gframework/http"
)

type Router struct {
	RouteCollection
	Server *http.Server
}

type RouterOptions struct {
	Port int
	SSL  bool
}

func NewRouter(options *RouterOptions) *Router {
	server := &http.Server{}
	router := NewRouterFromServer(server)
	//router.Routes = []*Route{}
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

func (r *Router) FindRouteByPath(path string) (*Route, _http.UrlParams) {

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
	println("Servidor escuchando en el puerto " + r.Server.Addr)
	log.Fatal(r.Server.ListenAndServe())
}
