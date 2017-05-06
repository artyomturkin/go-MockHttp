package main

import (
	"net/http"

	"strings"

	"gopkg.in/yaml.v2"
)

type Response struct {
	Code    int
	Headers map[string][]string
	Body    string
}

type Selector struct {
	In    string
	Key   string
	Value string
}

type Entry struct {
	Select   Selector
	Response Response
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SelectResponse(responses map[string]map[string][]Entry, r *http.Request) Response {
	response := Response{}
	for _, entry := range responses[r.URL.Path][r.Method] {
		if strings.ToLower(entry.Select.In) == "header" {
			if contains(r.Header[entry.Select.Key], entry.Select.Value) {
				return entry.Response
			}
		} else if strings.ToLower(entry.Select.In) == "query" {
			if contains(r.URL.Query()[entry.Select.Key], entry.Select.Value) {
				return entry.Response
			}
		} else {
			return entry.Response
		}
	}
	response.Code = 400
	return response
}

func Parse(responsesBytes []byte) map[string]map[string][]Entry {
	var Entries = make(map[string]map[string][]Entry)
	yaml.Unmarshal(responsesBytes, Entries)
	return Entries
}
