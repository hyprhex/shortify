package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Welcome to shortify"))
}

func linkView(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method not Allowed"))
		return
	}
	w.Write([]byte("view list of links"))
}

func linkCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not Allowed"))
		return
	}

	w.Write([]byte("Generate a short link"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/link/view", linkView)
	mux.HandleFunc("/link/create", linkCreate)

	log.Printf("Starting on server 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
