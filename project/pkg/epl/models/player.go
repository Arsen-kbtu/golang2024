package models

import (
	"context"
	"database/sql"
	"fmt"
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
func (m *PlayerModel) GetPlayers(firstname string, lastname string, age int, number int, nation string, position string, filters Filters) ([]*Player, error) {
	query := fmt.Sprintf(
		`
		SELECT  count(*) OVER(), *
		FROM players
		WHERE (STRPOS(LOWER(playerfirstname), LOWER($1)) > 0 OR $1= '')
		AND (STRPOS(LOWER(playerlastname), LOWER($2)) > 0 or $2 = '')
		AND ($3 = 0 OR playerage = $3)
		AND ($4 = 0 OR playernumber = $4)
		AND (STRPOS(LOWER(playernationality), LOWER($5)) > 0 or $5 = '')
		AND (STRPOS(LOWER(playerposition), LOWER($6)) > 0 or $6 = '')
		ORDER BY %s %s, playerid ASC
		LIMIT $7 OFFSET $8`, filters.sortColumn(), filters.sortDirection())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{firstname, lastname, age, number, nation, position, filters.limit(), filters.offset()}
	rows, err := m.DB.QueryContext(ctx, query, args...)
	fmt.Println(err)
	defer rows.Close()
	var players []*Player
	totalRecords := 0
	for rows.Next() {
		var player Player
		err := rows.Scan(&totalRecords, &player.PlayerID, &player.ClubID, &player.FirstName, &player.LastName, &player.Age, &player.Number, &player.Position, &player.Nation)
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
