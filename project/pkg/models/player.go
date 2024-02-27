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
			INSERT INTO players (clubid, firstname, lastname, age, number, position, nation)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id
			`
	args := []interface{}{player.ClubID, player.FirstName, player.LastName, player.Age, player.Number, player.Position, player.Nation}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&player.PlayerID)
}
func (m *PlayerModel) GetPlayer(id int) (*Player, error) {
	query := `
			SELECT id, clubid, firstname, lastname, age, number, position, nation
			FROM players
			WHERE id = $1
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
