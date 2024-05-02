-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Users (
    Id SERIAL PRIMARY KEY,
    Username VARCHAR(64) UNIQUE NOT NULL,
    Email VARCHAR(64) UNIQUE NOT NULL,
    Admin BOOL DEFAULT FALSE,
    Password CHAR(64) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
