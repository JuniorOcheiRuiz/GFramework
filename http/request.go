package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	RequestBase *http.Request
	URL         *url.URL
	Method      string
	Headers     http.Header
	Query       url.Values
	Body        io.ReadCloser
}

func NewRequest(requestBase *http.Request) *Request {
	request := &Request{RequestBase: requestBase}
	request.URL = requestBase.URL
	request.Method = requestBase.Method
	request.Body = requestBase.Body
	request.Query = requestBase.URL.Query()
	request.Headers = requestBase.Header

	err := parseBodyRequest(request)

	if err != nil {
		fmt.Println("Error parsing the body request.")
	}

	return request
}

func parseBodyRequest(request *Request) error {
	contentType := request.Headers.Get("Content-Type")
	var err error

	switch contentType {
	case "application/x-www-form-urlencoded":
		err = request.RequestBase.ParseForm()

	}

	return err
}
