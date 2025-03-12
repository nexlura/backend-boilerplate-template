-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
     "id" VARCHAR(255) PRIMARY KEY,
     "first_name" VARCHAR(100) NOT NULL,
     "last_name" VARCHAR(100) NOT NULL,
     "email" VARCHAR(100) UNIQUE NOT NULL,
     "password" VARCHAR(255) NOT NULL,
     "phone" VARCHAR(255) NOT NULL,
     "role_id" VARCHAR(50),
     "status" VARCHAR(50),
     "avatar" TEXT,
     "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
     "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

