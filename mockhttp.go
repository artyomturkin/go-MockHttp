package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"strconv"

	"fmt"

	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	//Setup and parse flags
	var mocksPath = flag.String("mocks", "", "path to a file with mock responses. Required")
	var port = flag.Int("port", 8080, "port to listen on.")
	var outPath = flag.String("out", "", "path to file to save incoming requests. If not specified outputs to stdout")
	flag.Parse()
	if *mocksPath == "" {
		panic("mocks flag must be set")
	}
	//Setup output
	var out io.Writer
	if *outPath == "" {
		out = os.Stdout
	} else {
		var err error
		out, err = os.OpenFile(*outPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			panic("could not open or create output file")
		}
	}
	//Parse responses
	responsesBytes, err := ioutil.ReadFile(*mocksPath)
	if err != nil {
		panic("could not read file with mock responses")
	}

	var entries = make(map[string]map[string][]Entry)
	err = yaml.Unmarshal(responsesBytes, entries)
	if err != nil {
		panic(err)
	}

	//Setup http server
	fmt.Printf("Listening on port %s", strconv.Itoa(*port))
	http.HandleFunc("/", MockHandler(entries, out))
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
