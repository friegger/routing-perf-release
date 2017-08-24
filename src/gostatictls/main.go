package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var data string

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/", HelloServer)
	port := os.Getenv("PORT")
	certBits := os.Getenv("APP_CERT")
	keyBits := os.Getenv("APP_KEY")
	s := os.Getenv("RESPONSE_SIZE")
	if s == "" {
		s = "1"
	}
	responseSize, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error parsing response size: %s", err)
	}

	err = ioutil.WriteFile("server.crt", []byte(certBits), 0400)
	if err != nil {
		log.Fatal("Error while writing cert: ", err)
	}
	err = ioutil.WriteFile("server.key", []byte(keyBits), 0400)
	if err != nil {
		log.Fatal("Error while writing key ", err)
	}

	fileBytes, err := ioutil.ReadFile("server.crt")
	log.Printf("cert contents %s", string(fileBytes))
	fileBytes, err = ioutil.ReadFile("server.key")
	log.Printf("key contents %s", string(fileBytes))

	data = strings.Repeat("Z", responseSize*1024)
	err = http.ListenAndServeTLS(fmt.Sprintf(":%s", port), "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
