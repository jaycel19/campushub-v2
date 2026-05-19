-- +goose Up
ALTER TABLE posts DROP COLUMN user_id;

ALTER TABLE posts
ADD COLUMN user_id uuid;



-- +goose Down
ALTER TABLE posts DROP COLUMN user_id;

ALTER TABLE posts
ADD COLUMN user_id int;