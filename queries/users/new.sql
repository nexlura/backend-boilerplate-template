INSERT INTO users (id, first_name, last_name, email, password, phone, role_id, status, avatar, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;
