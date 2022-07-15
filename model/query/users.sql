-- name: SelectUserID :one
select *
from users
where id_user = $1
and deleted is null;

-- name: UpdateUserOwnerValidation :one
update users
set owner_validation = true
where id_user = $1
RETURNING *;
;
