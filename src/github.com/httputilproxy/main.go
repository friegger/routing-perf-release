package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	reverseproxy := &httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL.Scheme = "http"
			request.URL.Host = "10.0.1.2:8080"
		},
	}

	server := &http.Server{Handler: reverseproxy, Addr: ":80"}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
