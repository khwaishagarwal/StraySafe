-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE incidents (
    Id SERIAL PRIMARY KEY,
    Uid INT NOT NULL,
    Latitude NUMERIC NOT NULL,
    Longitude NUMERIC NOT NULL,
    Title VARCHAR(128) NOT NULL,
    Description VARCHAR(1023),
    Image VARCHAR(64) NOT NULL,
    Resolved BOOLEAN NOT NULL DEFAULT FALSE,
    ResolverId INT DEFAULT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE incidents;
-- +goose StatementEnd
