SELECT "user".id, "user".email, "user".password_hash, "user".first_name, "user".last_name,
       COALESCE("user".avatar, '') as avatar,
       COALESCE("user".provider, '') as provider,
       COALESCE("user".phone, '') as phone,
       COALESCE("user".auth_token, '') as auth_token,
       "user".created_at, "user".updated_at
FROM users "user"
WHERE "user".email = $1 or "user".phone = $1;
