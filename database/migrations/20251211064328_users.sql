-- +goose Up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (), -- Google identity
    google_id text UNIQUE NOT NULL, -- Core identity
    email CITEXT NOT NULL UNIQUE,
    username CITEXT NOT NULL UNIQUE, -- at first gerated by the app
    avatar_url varchar(1024), -- from google accounts
    phone_number varchar(16) UNIQUE, -- Optional profile info
    phone_verified boolean NOT NULL DEFAULT FALSE, -- For now we can't verify phones
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    last_username_change timestamp DEFAULT NULL -- usename can be changed once every 14 days
);

CREATE INDEX IF NOT EXISTS users_email_idx ON users (email);

CREATE INDEX IF NOT EXISTS users_username_idx ON users (username);

CREATE INDEX IF NOT EXISTS users_google_id_idx ON users (google_id);

-- +goose Down
DROP TABLE IF EXISTS users;
