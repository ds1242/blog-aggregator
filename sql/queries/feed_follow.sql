-- name: FeedFollow :many
INSERT INTO feed_users(id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;