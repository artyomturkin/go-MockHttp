package main

import (
	"fmt"
	"io"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MockHandler(responses map[string]map[string][]Entry, out io.Writer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Output request
		requests := []Request{ParseRequest(r)}
		reqBuf, err := yaml.Marshal(requests)
		if err != nil {
			panic("failed to marshal request")
		}
		fmt.Fprintf(out, "%s", reqBuf)
		//Select response
		response := SelectResponse(responses, r)

		//Set headers
		for header, values := range response.Headers {
			for _, value := range values {
				w.Header().Add(header, value)
			}
		}
		w.WriteHeader(response.Code)
		fmt.Fprintf(w, "%s", response.Body)
	}
}
