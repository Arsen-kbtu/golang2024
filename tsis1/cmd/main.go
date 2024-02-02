package main

import (
	"net/http"
	"tsis1/pkg"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", pkg.WelcomeMessage).Methods("GET")
	router.HandleFunc("/health", pkg.HealthCheck).Methods("GET")
	router.HandleFunc("/arsenal", pkg.ShowTeam).Methods("GET")
	router.HandleFunc("/arsenal/{number}", pkg.PlayerByNum).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
