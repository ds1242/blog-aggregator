-- name: FeedFollow :one
INSERT INTO feed_users(id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UnfollowFeed :exec
DELETE FROM feed_users
WHERE feed_id = $1;
