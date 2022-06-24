-- +goose Up
-- +goose StatementBegin
CREATE TABLE usergroups (
    user_id         UUID NOT NULL REFERENCES users(id),
    group_id        UUID NOT NULL REFERENCES groups(id),
    CONSTRAINT usergroup_pk PRIMARY KEY (user_id, group_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE usergroups;
-- +goose StatementEnd
