-- name: GetFeeds :many
SELECT *
FROM feeds
ORDER BY updated_at;


-- name: AddToFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;