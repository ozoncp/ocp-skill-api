-- +goose Up
-- +goose StatementBegin
CREATE TABLE skills
(
    id          SERIAL     PRIMARY KEY,
    user_id     INTEGER     NOT NULL,
    name        TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE skills;
-- +goose StatementEnd
