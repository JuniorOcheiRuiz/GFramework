package routing

import (
	"fmt"
	"net/http"
)

type RequestHandler struct {
	router *Router
}

func NewRequestHandler(router *Router) *RequestHandler {
	return &RequestHandler{router: router}
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	route, params := h.router.FindRouteByPath(ctx.Request.Url.Path)
	routeName := "N/A"

	if route != nil {
		routeName = route.Name
		if route.HasMethod(ctx.Request.Method) {
			ctx.Params = params
			route.Handler(ctx)
		} else {
			ctx.Response.SetStatusCode(http.StatusMethodNotAllowed)
			_, _ = ctx.Response.WriteString("HTTP Method not allowed.")
		}
	} else {
		http.NotFound(w, r)
	}

	fmt.Printf("%s %s status:%d route:%s \n", ctx.Request.Method, ctx.Request.Url.Path, ctx.Response.GetStatusCode(), routeName)
}
