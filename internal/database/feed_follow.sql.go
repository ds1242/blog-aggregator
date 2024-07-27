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

const getUserFeed = `-- name: GetUserFeed :many
SELECT id, created_at, updated_at, feed_id, user_id 
FROM feed_users
WHERE user_id = $1
`

func (q *Queries) GetUserFeed(ctx context.Context, userID uuid.UUID) ([]FeedUser, error) {
	rows, err := q.db.QueryContext(ctx, getUserFeed, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedUser
	for rows.Next() {
		var i FeedUser
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FeedID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unfollowFeed = `-- name: UnfollowFeed :exec
DELETE FROM feed_users
WHERE id = $1
AND user_id = $2
`

type UnfollowFeedParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error {
	_, err := q.db.ExecContext(ctx, unfollowFeed, arg.ID, arg.UserID)
	return err
}
