-- +goose Up
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    -- Generate a random 64-character hexadecimal string as the API key
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
ALTER TABLE users
DROP COLUMN api_key;