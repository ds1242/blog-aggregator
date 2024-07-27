-- +goose Up
CREATE TABLE feed_users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL,
    user_id UUID NOT NULL,
    UNIQUE(feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_users;