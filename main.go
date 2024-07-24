package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	echoString := os.Getenv("ECHO_STRING")
	listenPort := os.Getenv("LISTEN_PORT")
	if echoString == "" {
		echoString = "Hello World"
	}
	if listenPort == "" {
		listenPort = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, echoString)
	})

	fmt.Println("Listening on port " + listenPort + " and answering with string " + echoString)
	http.ListenAndServe(":"+listenPort, nil)
}
