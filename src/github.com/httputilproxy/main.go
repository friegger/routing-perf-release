package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	reverseproxy := &httputil.ReverseProxy{
		Transport: transport,
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

var transport = &http.Transport{
	Dial: func(network, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(network, addr, 5*time.Second)
		if err != nil {
			return conn, err
		}
		err = conn.SetDeadline(time.Now().Add(900 * time.Second))
		return conn, err
	},
	DisableKeepAlives:  true,
	DisableCompression: true,
}
