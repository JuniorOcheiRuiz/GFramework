package http

import (
	"net/http"
)

type Response struct {
	ResponseBase http.ResponseWriter
}

func NewResponse(responseBase http.ResponseWriter) *Response {
	response := &Response{ResponseBase: responseBase}

	return response
}

func (r *Response) Write(value []byte) {
	_, _ = r.ResponseBase.Write(value)
}

func (r *Response) WriteString(value string) {
	r.Write([]byte(value))
}
