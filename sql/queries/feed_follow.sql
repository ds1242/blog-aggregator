-- name: FeedFollow :one
INSERT INTO feed_users(id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UnfollowFeed :exec
DELETE FROM feed_users
WHERE id = $1
AND user_id = $2;

-- name: GetUserFeed :many
SELECT * 
FROM feed_users
WHERE user_id = $1;
