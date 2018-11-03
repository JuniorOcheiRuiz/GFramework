package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	ResponseBase http.ResponseWriter
	statusCode   int
}

func NewResponse(responseBase http.ResponseWriter) *Response {
	response := &Response{ResponseBase: responseBase}
	response.statusCode = http.StatusOK
	return response
}

func (r *Response) Headers() http.Header {
	return r.ResponseBase.Header()
}

func (r *Response) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
	r.ResponseBase.WriteHeader(statusCode)
}

func (r *Response) GetStatusCode() int {
	return r.statusCode
}

func (r *Response) Write(value []byte) (int, error) {
	return r.ResponseBase.Write(value)
}

func (r *Response) WriteString(value string) (int, error) {
	return io.WriteString(r.ResponseBase, value)
}

func (r *Response) WriteJSON(v interface{}) (int, error) {
	json, err := json.Marshal(v)

	if err != nil {
		return 0, err
	}

	r.Headers().Set("Content-Type", "application/json")
	return r.Write(json)
}
