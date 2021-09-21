package main

import (
	"net/http"
	"os"
)

func main() {
	handlerFunc := http.HandlerFunc(handleRequest)
	http.Handle("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ipEnv := os.Getenv("AGENT-HOST")
	hostname, _ := os.Hostname()
	w.WriteHeader(200)

	if len(ipEnv) > 0 {
		w.Write([]byte("Hello World !!! in: " + ipEnv))
	} else {
		w.Write([]byte("Hello World !!! in: " + hostname))
	}
}
