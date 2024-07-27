// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const feedFollow = `-- name: FeedFollow :one
INSERT INTO feed_users(id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, feed_id, user_id
`

type FeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FeedID    uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) FeedFollow(ctx context.Context, arg FeedFollowParams) (FeedUser, error) {
	row := q.db.QueryRowContext(ctx, feedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FeedID,
		arg.UserID,
	)
	var i FeedUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FeedID,
		&i.UserID,
	)
	return i, err
}

const unfollowFeed = `-- name: UnfollowFeed :exec
DELETE FROM feed_users
WHERE feed_id = $1
`

func (q *Queries) UnfollowFeed(ctx context.Context, feedID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, unfollowFeed, feedID)
	return err
}
