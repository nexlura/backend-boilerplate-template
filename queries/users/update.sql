UPDATE users "user"
SET first_name = COALESCE(NULLIF($2,''), "user".first_name),
    last_name = COALESCE(NULLIF($3,''), "user".last_name),
    email = COALESCE(NULLIF($4,''), "user".email),
    password = COALESCE(NULLIF($5,''), "user".password),
    phone = COALESCE(NULLIF($6,''), "user".phone),
    role_id = COALESCE(NULLIF($7,''), "user".role_id),
    status = COALESCE(NULLIF($8,''), "user".status),
    avatar = COALESCE(NULLIF($9,''), "user".avatar),
    updated_at = CURRENT_TIMESTAMP
WHERE "user".id = $1;



