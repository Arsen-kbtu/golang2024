package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Arsenal Men Team!\n")
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is Arsenal Men Team!\nMade by Arsen\n/arsenal shows current team members\n/arsenal/any_number shows information about player with that number")
}
func ShowTeam(w http.ResponseWriter, r *http.Request) {
	team := GetTeam()
	// fmt.Fprintf(w, "Shirt Number  Name  Surname\n")
	// for i := 0; i < len(team); i++ {
	// 	fmt.Fprintf(w, "%d  %s %s  \n", team[i].Number, team[i].FirstName, team[i].LastName)
	// }
	json.NewEncoder(w).Encode(team)
}
func PlayerByNum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["number"])
	if err != nil {
		return
	}
	player, found := FindByNum(num)
	if found {
		// fmt.Fprintf(w, "%d\nName: %s\nSurname: %s\nAge: %d\nPosition: %s\n", player.Number, player.FirstName, player.LastName, player.Age, player.Position)
		json.NewEncoder(w).Encode(player)
	} else {
		fmt.Fprintf(w, "There is no player with such shirt number!")
	}

}
