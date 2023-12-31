package main

import (
	"base3/gee"
	"fmt"
	"net/http"
)

func main() {

	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q = %q]\n", k, v)
		}
	})
	r.Run(":9999")
}
