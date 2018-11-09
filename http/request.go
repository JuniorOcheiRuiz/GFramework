package http

import (
	"encoding/json"
	"fmt"
	"io"
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
	Params        UrlParams
	UploadedFiles map[string][]*UploadedFile
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

	parseMultipartForm(request)

	return request
}

func parseMultipartForm(request *Request) {
	for index, file := range request.RequestBase.MultipartForm.File {
		var uploadedFiles []*UploadedFile

		for _, fileHeader := range file {
			uploadedFile := NewUploadFile(fileHeader)
			uploadedFiles = append(uploadedFiles, uploadedFile)
		}

		request.UploadedFiles[index] = uploadedFiles
	}
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

func (r *Request) ParseJSON(value interface{}) error {
	decoder := json.NewDecoder(r.Body)

	return decoder.Decode(value)
}
