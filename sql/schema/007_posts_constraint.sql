-- +goose Up
ALTER TABLE posts
ADD UNIQUE (url);


-- +goose Down
ALTER TABLE posts
DROP CONSTRAINT url;