package main

import (
	"encoding/json"
	"net/http"
	pkg "project/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func (app *application) getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid player number")
		return
	}

	player, err := app.models.Players.GetPlayer(number)

	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "Player not found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, player)
}
func (app *application) getClubHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid club ID")
		return
	}

	club, err := app.models.Clubs.GetClub(id)

	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "Club not found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, club)
}
func (app *application) createClubHandler(w http.ResponseWriter, r *http.Request) {
	var club pkg.Club
	err := json.NewDecoder(r.Body).Decode(&club)

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if club.ClubName == "" || club.ClubCity == "" {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = app.models.Clubs.InsertClub(&club)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, club)
}
