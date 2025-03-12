-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, first_name, last_name, email, password, phone, role_id, status)
    VALUES ('1','demo', 'user', 'demo@user.com', '12345678', '', 'system_admin', 'active');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id='1';
-- +goose StatementEnd

