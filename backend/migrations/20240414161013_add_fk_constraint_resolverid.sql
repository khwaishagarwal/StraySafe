-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE incidents
    ADD CONSTRAINT fk_resolverid FOREIGN KEY(ResolverId) REFERENCES Ngos(Id)
        ON DELETE SET NULL
        ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP CONSTRAINT fk_resolverid;
-- +goose StatementEnd
