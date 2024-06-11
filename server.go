package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/new-endpoint", handleNewEndpoint)

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(w, "Hello World!")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(w, "ok")
	fmt.Println("Handle health endpoint triggered")
}

func writeResponse(w http.ResponseWriter, responseString string) {
	response := []byte(responseString)

	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

func handleNewEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	fmt.Println("Handle new endpoint triggered")
	writeResponse(w, "new endpoint")
}
