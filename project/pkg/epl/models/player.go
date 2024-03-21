package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

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
type PlayerModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m *PlayerModel) InsertPlayer(player *Player) error {
	query := `
			INSERT INTO players (playerclubid, playerfirstname, playerlastname, playerage, playernumber, playerposition, playernationality)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING playerid
			`
	args := []interface{}{player.ClubID, player.FirstName, player.LastName, player.Age, player.Number, player.Position, player.Nation}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&player.PlayerID)
}
func (m *PlayerModel) GetPlayer(id int) (*Player, error) {
	query := `
			SELECT playerid, playerclubid,playerfirstname, playerlastname, playerage, playernumber, playerposition, playernationality
			FROM players
			WHERE playerid = $1
			`
	var player Player
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&player.PlayerID, &player.ClubID, &player.FirstName, &player.LastName, &player.Age, &player.Number, &player.Position, &player.Nation)
	if err != nil {
		return nil, err
	}
	return &player, nil
}
func (m *PlayerModel) GetPlayers() ([]*Player, error) {
	query := `
			SELECT playerid, playerclubid,playerfirstname, playerlastname, playerage, playernumber, playerposition, playernationality
			FROM players
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []*Player
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.PlayerID, &player.ClubID, &player.FirstName, &player.LastName, &player.Age, &player.Number, &player.Position, &player.Nation)
		if err != nil {
			return nil, err
		}
		players = append(players, &player)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return players, nil
}
func (m *PlayerModel) UpdatePlayer(player *Player) error {
	query := `
			UPDATE players
			SET playerclubid = $1, playerfirstname = $2, playerlastname = $3, playerage = $4, playernumber = $5, playerposition = $6, playernationality = $7
			WHERE playerid = $8
			`
	args := []interface{}{player.ClubID, player.FirstName, player.LastName, player.Age, player.Number, player.Position, player.Nation, player.PlayerID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}
func (m *PlayerModel) DeletePlayer(id int) error {
	query := `
			DELETE FROM players
			WHERE playerid = $1
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
