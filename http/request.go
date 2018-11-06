package http

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Request struct {
	RequestBase   *http.Request
	URL           *url.URL
	Method        string
	Headers       http.Header
	Query         url.Values
	Body          io.ReadCloser
	Form          url.Values
	MultipartForm *multipart.Form
}

func NewRequest(r *http.Request) *Request {
	request := &Request{RequestBase: r}
	request.URL = r.URL
	request.Method = r.Method
	request.Body = r.Body
	request.Query = r.URL.Query()
	request.Headers = r.Header

	err := parseBodyRequest(request)

	if err != nil {
		fmt.Println("Error parsing the body request.", err)
	}

	request.Form = r.Form
	request.MultipartForm = r.MultipartForm

	return request
}

func parseBodyRequest(request *Request) error {
	contentType := request.Headers.Get("Content-Type")
	var err error

	switch contentType {
	case "application/x-www-form-urlencoded":
		err = request.RequestBase.ParseForm()
	case "multipart/form-data":
		err = request.RequestBase.ParseMultipartForm(0)
	}

	return err
}
