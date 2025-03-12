UPDATE users "user"
SET first_name = COALESCE(NULLIF($1,''), "user".first_name),
    last_name = COALESCE(NULLIF($2,''), "user".last_name),
    email = COALESCE(NULLIF($3,''), "user".email),
    password_hash = COALESCE(NULLIF($4,''), "user".password_hash),
    phone = COALESCE(NULLIF($5,''), "user".phone),
    address = COALESCE(NULLIF($6,''), "user".address),
    country = COALESCE(NULLIF($7,''), "user".country),
    account_type = COALESCE(NULLIF($8,''), "user".account_type),
    status = COALESCE(NULLIF($9,''), "user".status),
    avatar = COALESCE(NULLIF($10,''), "user".avatar),
    auth_token = COALESCE(NULLIF($11,''), "user".auth_token),
    updated_at = CURRENT_TIMESTAMP
WHERE "user".id = $12;



