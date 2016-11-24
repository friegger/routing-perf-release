package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/vulcand/oxy/forward"
)

type nilLogger struct{}

func (n nilLogger) Write(p []byte) (int, error) {
	return 0, nil
}

func main() {
	logrus.SetOutput(nilLogger{})
	fwd, _ := forward.New()

	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// let us forward this request to another server
		req.URL.Scheme = "http"
		req.URL.Host = "10.0.1.2:8080"
		fwd.ServeHTTP(w, req)
	})

	server := &http.Server{Handler: redirect, Addr: ":80"}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
