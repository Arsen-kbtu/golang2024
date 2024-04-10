package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	pkg "project/pkg/epl/models"
	"project/pkg/epl/validator"

	"strconv"

	"github.com/gorilla/mux"
)

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
	var input struct {
		ClubName string
		Clubcity string
		pkg.Filters
	}
	// Initialize a new Validator instance.
	v := validator.New()
	// Call r.URL.Query() to get the url.Values map containing the query string data.
	qs := r.URL.Query()
	// Use our helpers to extract the title and genres query string values, falling back
	// to defaults of an empty string and an empty slice respectively if they are not
	// provided by the client.
	input.ClubName = app.readString(qs, "clubname", "")
	input.Clubcity = app.readString(qs, "clubcity", "")
	// Get the page and page_size query string values as integers. Notice that we set
	// the default page value to 1 and default page_size to 20, and that we pass the
	// validator instance as the final argument here.
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	// Extract the sort query string value, falling back to "id" if it is not provided
	// by the client (which will imply a ascending sort on movie ID).
	input.Filters.Sort = app.readString(qs, "sort", "clubid")
	input.Filters.SortSafelist = []string{"clubid", "clubname", "clubcity", "leagueplacement", "-clubid", "-clubname", "-clubcity", "-leagueplacement"}
	// Check the Validator instance for any errors and use the failedValidationResponse()
	// helper to send the client a response if necessary.
	if pkg.ValidateFilters(v, input.Filters); !v.Valid() {
		//app.failedValidationResponse(w, r, v.Errors)
		return
	}
	if !v.Valid() {
		//app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Dump the contents of the input struct in a HTTP response.
	//fmt.Fprintf(w, "%+v\n", input)

	clubs, metadata, err := app.models.Clubs.GetClubs(input.ClubName, input.Clubcity, input.Filters)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}
	fmt.Println(metadata)
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
