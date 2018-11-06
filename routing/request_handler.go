package routing

import (
	"fmt"
	_http "gframework/http"
	"net/http"
)

type RequestHandler struct {
	router *Router
}

func NewRequestHandler(router *Router) *RequestHandler {
	return &RequestHandler{router: router}
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := _http.NewRequest(r)
	response := _http.NewResponse(w)

	route, params := h.router.FindRouteByPath(request.URL.Path)
	routeName := "N/A"

	if route != nil {
		routeName = route.Name
		if route.HasMethod(request.Method) {
			request.Params = params
			route.Handler(request, response)
		} else {
			response.SetStatusCode(http.StatusMethodNotAllowed)
			_, _ = response.WriteString("HTTP Method not allowed.")
		}
	} else {
		http.NotFound(w, r)
	}

	fmt.Printf("%s %s status:%d route:%s \n", request.Method, request.URL.Path, response.GetStatusCode(), routeName)
}
