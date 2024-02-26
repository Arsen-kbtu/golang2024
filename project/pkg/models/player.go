package models

type Player struct {
	PlayerID  int    `json:"playerid"`
	ClubID    int    `json:"clubid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Number    int    `json:"number"`
	Position  string `json:"position"`
	Nation    string `json:"nation"`
}
