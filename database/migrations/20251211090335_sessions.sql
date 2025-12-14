-- +goose Up
CREATE TABLE IF NOT EXISTS sessions (
    token text PRIMARY KEY,
    data bytea NOT NULL,
    expiry timestamptz NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS sessions;
