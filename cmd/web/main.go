package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/link/views", linkViews)
	mux.HandleFunc("/link/view", linkView)
	mux.HandleFunc("/link/create", linkCreate)

	log.Printf("Starting on server 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
