-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS groups (
    id          UUID PRIMARY KEY,
    title       TEXT NOT NULL,
    group_type  TEXT NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    CONSTRAINT group_unique_pk UNIQUE (title, group_type)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE groups;
-- +goose StatementEnd
