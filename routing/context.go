package routing

import (
	_http "gframework/http"
	"net/http"
)

type Context struct {
	Request  *_http.Request
	Response *_http.Response
	Params   RouteParams
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{Request: _http.NewRequest(r), Response: _http.NewResponse(w)}
}
