package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to shortify"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("Starting on server 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
