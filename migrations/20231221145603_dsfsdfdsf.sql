-- +goose Up
SELECT 'up SQL query';

-- +goose Down
SELECT 'down SQL query';
-- bin/goose -dir migrations postgres "host=localhost port=5432 user=user password=user123 sslmode=disable dbname=db_name" up -v