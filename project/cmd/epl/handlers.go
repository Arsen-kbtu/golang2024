package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	pkg "project/pkg/epl/models"

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

// player handlers
func (app *application) getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["playerId"])
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid player number")
		return
	}

	player, err := app.models.Players.GetPlayer(id)

	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "Player not found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, player)
}
func (app *application) getPlayersHandler(w http.ResponseWriter, r *http.Request) {
	players, err := app.models.Players.GetPlayers()

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, players)

}
func (app *application) createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	var player pkg.Player
	err := json.NewDecoder(r.Body).Decode(&player)

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if player.FirstName == "" || player.LastName == "" {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = app.models.Players.InsertPlayer(&player)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, player)
}

func (app *application) updatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	var player pkg.Player
	err := json.NewDecoder(r.Body).Decode(&player)

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if player.FirstName == "" || player.LastName == "" {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = app.models.Players.UpdatePlayer(&player)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, player)
}

func (app *application) deletePlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["playerId"])

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid player ID")
		return
	}

	err = app.models.Players.DeletePlayer(id)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusNoContent, nil)
}

// Club Handlers
func (app *application) getClubsHandler(w http.ResponseWriter, r *http.Request) {
	clubs, err := app.models.Clubs.GetClubs()

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, clubs)
}
func (app *application) getClubHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["clubId"])
	fmt.Println(id)
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

func (app *application) updateClubHandler(w http.ResponseWriter, r *http.Request) {
	var club pkg.Club
	err := json.NewDecoder(r.Body).Decode(&club)

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid club ID")
		return
	}

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if club.ClubName == "" || club.ClubCity == "" {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = app.models.Clubs.UpdateClub(&club)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, club)
}

func (app *application) deleteClubHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["clubId"])

	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid club ID")
		return
	}

	err = app.models.Clubs.DeleteClub(id)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusNoContent, nil)
}
