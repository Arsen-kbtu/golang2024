package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Club struct {
	ClubID       int    `json:"clubid"`
	ClubName     string `json:"clubname"`
	ClubCity     string `json:"clubcity"`
	LeaguePlace  int    `json:"position"`
	LeaguePoints int    `json:"points"`
}
type ClubModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *ClubModel) InsertClub(club *Club) error {
	query := `
			INSERT INTO clubs (clubname, clubcity, leagueplace, leaguepoints) 
			VALUES ($1, $2, $3, $4)
			RETURNING id
			`
	args := []interface{}{club.ClubName, club.ClubCity, club.LeaguePlace, club.LeaguePoints}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&club.ClubID)
}
func (m *ClubModel) GetClubs() ([]*Club, error) {
	return nil, nil
}
func (m *ClubModel) GetClub(id int) (*Club, error) {
	return nil, nil
}
func (m *ClubModel) UpdateClub(id int, name, city string) error {
	return nil
}
