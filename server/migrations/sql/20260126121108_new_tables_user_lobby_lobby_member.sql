-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id   uuid PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE lobbies
(
    id   uuid PRIMARY KEY,
    name text NOT NULL,
    owner uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE lobby_members
(
    lobby_id uuid NOT NULL REFERENCES lobbies (id) ON DELETE CASCADE,
    user_id  uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (lobby_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lobby_members;
DROP TABLE IF EXISTS lobbies;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
