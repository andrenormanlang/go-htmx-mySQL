-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  content TEXT NOT NULL, -- HTML content
  link TEXT NOT NULL, -- URL
  title TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd