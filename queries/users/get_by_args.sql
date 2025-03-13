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
FROM users "user"
WHERE "user".id = $1 OR "user".email = $1 OR "user".phone = $1;
