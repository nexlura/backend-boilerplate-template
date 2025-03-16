-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles (
    "id" VARCHAR(255) PRIMARY KEY,
    "title" VARCHAR(100) NOT NULL,
    "provisioned_by" VARCHAR(100),
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
