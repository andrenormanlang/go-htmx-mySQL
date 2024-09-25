-- +goose Up
-- +goose StatementBegin
UPDATE posts SET excerpt = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam.';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
UPDATE posts SET excerpt = NULL;
-- +goose StatementEnd
