package main

import (
	"io/ioutil"
	"net/http"
)

type Request struct {
	URL     string
	Method  string
	Headers map[string][]string
	Body    string
}

func ParseRequest(r *http.Request) Request {
	request := Request{}
	request.URL = r.URL.String()
	request.Headers = r.Header
	request.Method = r.Method

	bodyBuf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("failed to read request.Body")
	}
	request.Body = string(bodyBuf)

	return request
}
