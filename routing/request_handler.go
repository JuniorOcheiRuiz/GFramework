package routing

import "net/http"

type RequestHandler struct {
	router *Router
}

func NewRequestHandler(router *Router) *RequestHandler {
	return &RequestHandler{router: router}
}

func (h *RequestHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}
