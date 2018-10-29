package routing

import "net/http"

type RequestHandler struct {
	router *Router
}

func NewRequestHandler(router *Router) *RequestHandler {
	return &RequestHandler{router: router}
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	route, params := h.router.FindRouteByPath(ctx.Request.Url.Path)

	if route != nil {
		ctx.Params = params
		route.Handler(ctx)
	} else {
		http.NotFound(w, r)
	}
}
