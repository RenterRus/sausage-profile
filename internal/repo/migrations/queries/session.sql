-- name: GetRefreshToken :one
select refresh_hash, 
(expired_at <= now()) as is_expired, 
exists(select 1 from blacklist_refresh where refresh_hash = (select refresh_hash from users u where u.user_login = sqlc.narg('login'))) as block 
from users u where u.user_login = sqlc.narg('login');

-- name: SetBlockRefresh :exec
insert into blacklist_refresh (refresh_hash) values (sqlc.narg('refresh_hash'));

-- name: SetRefreshHash :exec
update users 
set refresh_hash = sqlc.narg('refresh_hash'),
expired_at = sqlc.narg('expired_at')
where user_login = sqlc.narg('login');
