package models

type Club struct {
	ClubID       int    `json:"clubid"`
	ClubName     string `json:"clubname"`
	ClubCity     string `json:"clubcity"`
	LeaguePlace  int    `json:"position"`
	LeaguePoints int    `json:"points"`
}
