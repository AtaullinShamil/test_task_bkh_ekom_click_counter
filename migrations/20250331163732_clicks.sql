-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS counter;

CREATE TABLE counter.clicks (
    id SERIAL PRIMARY KEY,
    banner_id INT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE INDEX idx_clicks_banner_id ON counter.clicks(banner_id);

CREATE INDEX idx_clicks_timestamp ON counter.clicks(timestamp);

CREATE INDEX idx_clicks_banner_id_timestamp ON counter.clicks(banner_id, timestamp);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS counter CASCADE;
-- +goose StatementEnd
