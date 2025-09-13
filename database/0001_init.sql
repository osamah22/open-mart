
-- +goose Up

create extension if not exists "uuid-ossp"; --an extension responsible for generating uuid

CREATE TABLE IF NOT EXISTS categories (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_id  uuid, -- for sub-directories
    name VARCHAR(255) UNIQUE NOT NULL,
);

CREATE TABLE IF NOT EXISTS cities (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar(128) UNIQUE NOT NULL,
    name_ar varchar(128) UNIQUE NOT NULL 
);

-- +goose Down
DROP TABLE IF NOT EXISTS cities;
DROP TABLE IF NOT EXISTS categories;
