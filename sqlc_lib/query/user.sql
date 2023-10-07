-- name: CreateAdminUser :one
insert into users (
    name,
    email,
    contact,
    password,
    user_type,
    is_account_active
) values (
    $1,$2,$3,$4,$5,$6
) returning *;

-- name: GetUserById :one
select * from
users where id = $1
limit 1;

-- name: ActiveDeactiveUserAccount :one
update users 
set is_account_active=$2,
updated_at = CURRENT_TIMESTAMP
where id = $1
returning *
;