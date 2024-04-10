CREATE TABLE IF NOT EXISTS clubs (
                                     clubId SERIAL PRIMARY KEY,
                                     clubName VARCHAR(255) NOT NULL,
    clubCity VARCHAR(255) NOT NULL,
    leaguePlacement INT NOT NULL,
    leaguePoints INT NOT NULL
    );

CREATE TABLE IF NOT EXISTS players (
                                       playerId SERIAL PRIMARY KEY,
                                       playerClubId INT NOT NULL,
                                       playerFirstName VARCHAR(255) NOT NULL,
    playerLastName VARCHAR(255) NOT NULL,
    playerAge INT NOT NULL,
    playerNumber INT NOT NULL,
    playerPosition VARCHAR(255) NOT NULL,
    playerNationality VARCHAR(255) NOT NULL,
    FOREIGN KEY (playerClubId) REFERENCES clubs(clubId)
    );

CREATE TABLE IF NOT EXISTS model (
                                     modelId SERIAL PRIMARY KEY,
                                     clubId INT NOT NULL,
                                     playerId INT NOT NULL,
                                     FOREIGN KEY (clubId) REFERENCES clubs(clubId),
    FOREIGN KEY (playerId) REFERENCES players(playerId)
    );
CREATE TABLE IF NOT EXISTS Users
(
    user_id  SERIAL PRIMARY KEY,
    username VARCHAR(50)  NOT NULL,
    password VARCHAR(100) NOT NULL
    );