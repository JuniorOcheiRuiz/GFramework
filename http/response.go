package http

import (
	"bytes"
	"log"
	"net/http"
)

type Response struct {
	ResponseBase http.ResponseWriter
	StatusCode   int
	Body         *bytes.Buffer
}

func NewResponse(responseBase http.ResponseWriter) *Response {
	response := &Response{ResponseBase: responseBase}
	response.StatusCode = 200
	response.Body = bytes.NewBuffer([]byte{})
	return response
}

func (r *Response) Write(value []byte) {
	_, err := r.Body.Write(value)

	if err != nil {
		log.Print("Error writing in the body of response.", err)
	}
}

func (r *Response) WriteString(value string) {
	r.Write([]byte(value))
}
