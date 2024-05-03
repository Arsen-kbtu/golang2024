CREATE EXTENSION IF NOT EXISTS citext;
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
CREATE TABLE IF NOT EXISTS users (
                                     id bigserial PRIMARY KEY,
                                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                     name text NOT NULL,
                                     email citext UNIQUE NOT NULL,
                                     password_hash bytea NOT NULL,
                                     activated bool NOT NULL,
                                     version integer NOT NULL DEFAULT 1
);
CREATE TABLE IF NOT EXISTS tokens (
                                      hash bytea PRIMARY KEY,
                                      user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
                                      expiry timestamp(0) with time zone NOT NULL,
                                      scope text NOT NULL
);
CREATE TABLE IF NOT EXISTS permissions (
                                           id bigserial PRIMARY KEY,
                                           code text NOT NULL
);
CREATE TABLE IF NOT EXISTS users_permissions (
                                                 user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
                                                 permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
                                                 PRIMARY KEY (user_id, permission_id)
);
INSERT INTO permissions (code)
VALUES
('user.create'),
('user.read'),
('user.update'),
('user.delete'),
('club.create'),
('club.read'),
('club.update'),
('club.delete'),
('player.create'),
('player.read'),
('player.update'),
('player.delete');