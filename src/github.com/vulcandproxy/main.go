package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vulcand/oxy/forward"
)

func main() {
	fwd, _ := forward.New()

	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// let us forward this request to another server
		req.URL.Scheme = "http"
		req.URL.Host = "10.0.1.2:8080"
		fwd.ServeHTTP(w, req)
	})

	//	reverseproxy := &httputil.ReverseProxy{
	//		Transport: transport,
	//		Director: func(request *http.Request) {
	//			request.URL.Scheme = "http"
	//			request.URL.Host = "10.0.1.2:8080"
	//		},
	//	}

	server := &http.Server{Handler: redirect, Addr: ":80"}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
