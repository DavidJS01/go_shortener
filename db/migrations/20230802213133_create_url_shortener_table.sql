-- +goose Up
-- +goose StatementBegin
create table if not exists url_shortener (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     long_url TEXT UNIQUE,
     shortened_url TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table url_shortener;
-- +goose StatementEnd
