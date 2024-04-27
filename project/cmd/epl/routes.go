package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Menu Singleton
	// Create a new menu
	v1.HandleFunc("/clubs", app.createClubHandler).Methods("POST")
	// Get a specific menu
	v1.HandleFunc("/clubs", app.getClubsHandler).Methods("GET")
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.requireActivatedUser(app.getClubHandler)).Methods("GET")
	// Update a specific menu
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.requirePermission("club.update", app.updateClubHandler)).Methods("PUT")
	// Delete a specific menu
	v1.HandleFunc("/clubs/{clubId:[0-9]+}", app.requirePermission("club.delete", app.deleteClubHandler)).Methods("DELETE")

	v1.HandleFunc("/players", app.createPlayerHandler).Methods("POST")
	v1.HandleFunc("/players", app.getPlayersHandler).Methods("GET")
	v1.HandleFunc("/players/{playerId:[0-9]+}", app.requireActivatedUser(app.getPlayerHandler)).Methods("GET")
	v1.HandleFunc("/players/{playerId:[0-9]+}", app.updatePlayerHandler).Methods("PUT")
	v1.HandleFunc("/players/{playerId:[0-9]+}", app.deletePlayerHandler).Methods("DELETE")

	v1.HandleFunc("/users", app.registerUserHandler).Methods("POST")
	v1.HandleFunc("/users/activated", app.activateUserHandler).Methods("PUT")
	v1.HandleFunc("/tokens/authentication", app.createAuthenticationTokenHandler).Methods("POST")

	log.Printf("Starting server on %d\n", app.config.port)
	//err := http.ListenAndServe(app.config.port, r)

	return app.authenticate(r)
}
