SELECT
    id,
    first_name,
    last_name,
    email,
    password,
    COALESCE(phone, '') as phone,
    COALESCE(role_id, '') as role_id,
    COALESCE(status, '') as status,
    COALESCE(avatar, '') as avatar,
    COALESCE(auth_token, '') as auth_token,
    created_at,
    updated_at
FROM users
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2;
