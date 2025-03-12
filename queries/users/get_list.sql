SELECT
    id,
    first_name,
    last_name,
    email,
    COALESCE(phone, '') as phone,
    COALESCE(address, '') as address,
    COALESCE(country, '') as country,
    COALESCE(account_type, '') as account_type,
    COALESCE(status, '') as status,
    COALESCE(avatar, '') as avatar,
    COALESCE(provider, '') as provider,
    COALESCE(auth_token, '') as auth_token,
    created_at,
    updated_at
FROM users
ORDER BY updated_at DESC
LIMIT $1 OFFSET $2;
