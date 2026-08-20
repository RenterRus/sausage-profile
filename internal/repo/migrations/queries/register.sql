-- name: Register :exec
insert into users (user_login, uuid, otp_hash, otp_link, created_at, last_sign_up_at) values (sqlc.narg('login'), gen_random_uuid(), sqlc.narg('hash'), sqlc.narg('link'), now(), now());

-- name: Confirmed :exec
update users set confirmed = true where user_login = sqlc.narg('login');

-- name: IsConfirmed :one
select confirmed from users where user_login = sqlc.narg('login');

-- name: Hash :one
select otp_hash from users where user_login = sqlc.narg('login');