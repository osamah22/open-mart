-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    category_id uuid NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

-- +goose Down
CREATE TABLE IF NOT EXISTS posts;
