package routing

import (
	"fmt"
	"log"
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
			ctx.Response.StatusCode = http.StatusMethodNotAllowed
			ctx.Response.WriteString("HTTP Method not allowed.")
		}
	} else {
		http.NotFound(w, r)
	}

	w.WriteHeader(ctx.Response.StatusCode)
	_, err := w.Write(ctx.Response.Body.Bytes())

	if err != nil {
		log.Print("Error writing in the body of response.", err)
	}

	fmt.Printf("%s %s status:%d route:%s \n", ctx.Request.Method, ctx.Request.Url.Path, ctx.Response.StatusCode, routeName)
}
