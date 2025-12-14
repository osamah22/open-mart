-- name: GetUserByGoogleID :one
SELECT *
FROM users
WHERE google_id = $1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (google_id, email, username, phone_number, avatar_url)
    VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUserAvatar :exec
UPDATE
    users
SET avatar_url = $2,
    updated_at = now()
WHERE id = $1;

-- name: UpdateUsername :one
UPDATE
    users
SET username = $2,
    updated_at = now(),
    last_username_change = now()
WHERE id = $1
RETURNING *;

-- name: UsernameExists :one
SELECT EXISTS (
        SELECT 1
        FROM users
        WHERE username = $1);
