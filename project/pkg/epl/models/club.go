package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Club struct {
	ClubID       int      `json:"clubid"`
	ClubName     string   `json:"clubname"`
	ClubCity     string   `json:"clubcity"`
	LeaguePlace  int      `json:"position"`
	LeaguePoints int      `json:"points"`
	Players      []Player `json:"players"`
}
type ClubModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *ClubModel) InsertClub(club *Club) error {
	query := `
			INSERT INTO clubs (clubname, clubcity, leagueplacement, leaguepoints) 
			VALUES ($1, $2, $3, $4)
			RETURNING clubid
			`
	args := []interface{}{club.ClubName, club.ClubCity, club.LeaguePlace, club.LeaguePoints}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&club.ClubID)
}
func (m *ClubModel) GetClubs() ([]*Club, error) {
	query := `
			SELECT clubid, clubname, clubcity, leagueplacement, leaguepoints
			FROM clubs
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clubs []*Club
	for rows.Next() {
		var club Club
		err := rows.Scan(&club.ClubID, &club.ClubName, &club.ClubCity, &club.LeaguePlace, &club.LeaguePoints)
		if err != nil {
			return nil, err
		}
		clubs = append(clubs, &club)

		playersQuery := `
        SELECT *
        FROM players
        WHERE playerclubid = $1
    `
		rows, err := m.DB.QueryContext(ctx, playersQuery, club.ClubID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Проходим по строкам результата и собираем игроков
		for rows.Next() {
			var player Player
			err := rows.Scan(&player.PlayerID, &player.ClubID, &player.FirstName, &player.LastName, &player.Age, &player.Number, &player.Position, &player.Nation)
			if err != nil {
				return nil, err
			}
			club.Players = append(club.Players, player)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return clubs, nil
}
func (m *ClubModel) GetClub(id int) (*Club, error) {
	query := `
        SELECT *
        FROM clubs
        WHERE clubid = $1
    `
	var club Club
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&club.ClubID, &club.ClubName, &club.ClubCity, &club.LeaguePlace, &club.LeaguePoints)
	if err != nil {
		return nil, err
	}

	// Теперь запросим список игроков для данного клуба
	playersQuery := `
        SELECT *
        FROM players
        WHERE playerclubid = $1
    `
	rows, err := m.DB.QueryContext(ctx, playersQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Проходим по строкам результата и собираем игроков
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.PlayerID, &player.ClubID, &player.FirstName, &player.LastName, &player.Age, &player.Number, &player.Position, &player.Nation)
		if err != nil {
			return nil, err
		}
		club.Players = append(club.Players, player)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &club, nil
}
func (m *ClubModel) UpdateClub(club *Club) error {
	query := `
			UPDATE clubs
			SET clubname = $1,
			clubcity = $2,
			leagueplacement = $3,
			leaguepoints = $4
			WHERE clubid = $5
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, query, club.ClubName, club.ClubCity, club.LeaguePlace, club.LeaguePoints, club.ClubID)
	return err
}
func (m *ClubModel) DeleteClub(id int) error {
	query := `
			DELETE FROM clubs
			WHERE clubid = $1
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
