package main

import (
	"net/http"

	"gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/").Methods("GET")
	router.HandleFunc("/health").Methods("GET")
	router.HandleFunc("/arsenal").Methods("GET")
	router.HandleFunc("/arsenal/{number}").Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
