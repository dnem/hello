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

	cr := os.Getenv("CONTEXT_ROOT")
	if len(cr) == 0 {
		cr = ""
	} else {
		if string(cr[0]) != "/" {
			cr = "/" + cr
		}
		if string(cr[len(cr)-1]) == "/" {
			cr = cr[:len(cr)-1]
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc(cr+"/", hello)
	mux.HandleFunc(cr+"/nonroot", nonroot)

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
