-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Ngos (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(128) NOT NULL UNIQUE,
    Email VARCHAR(128) NOT NULL UNIQUE,
    About VARCHAR(4096) NOT NULL,
    Latitude NUMERIC NOT NULL,
    Longitude NUMERIC NOT NULL,
    Password VARCHAR(60) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Ngo;
-- +goose StatementEnd
