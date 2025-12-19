-- name: GetPostById :many
SELECT *
FROM posts
WHERE id = $1;

-- name: ListPosts :many
SELECT *
FROM posts
ORDER BY created_at OFFSET sqlc.arg (page_number) * sqlc.arg (page_size)
LIMIT sqlc.arg (page_size);

-- name: CreatePost :one
INSERT INTO posts (title, content, user_id, category_id)
    VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePost :one
UPDATE
    posts
SET title = $2,
    content = $3,
    user_id = $4,
    category_id = $5
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
