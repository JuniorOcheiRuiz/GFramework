package http

import (
	"net/http"
	"net/url"
)

type Request struct {
	RequestBase *http.Request
	Url         *url.URL
	Method      string
}

func NewRequest(requestBase *http.Request) *Request {
	request := &Request{RequestBase: requestBase}
	request.Url = requestBase.URL
	request.Method = requestBase.Method

	return request
}
