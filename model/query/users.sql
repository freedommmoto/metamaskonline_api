-- name: SelectUserID :one
select *
from users
where id_user = $1
and deleted is null;
