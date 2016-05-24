package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/nonroot", nonroot)

	n := negroni.Classic()
	n.UseHandler(mux)
	hostString := fmt.Sprintf(":%s", port)
	n.Run(hostString)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from root!")
}

func nonroot(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from somplace that is not the root!")
}
