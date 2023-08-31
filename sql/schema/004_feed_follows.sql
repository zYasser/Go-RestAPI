-- +goose Up
CREATE TABLE feed_follow (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
    
);

-- +goose Down
DROP TABLE feeds;