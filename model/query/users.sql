-- name: SelectUserID :one
select *
from users
where id_user = $1
  and deleted is null;

-- name: SelectUserByLineUserID :one
select *
from users
where id_line = $1
  and deleted is null;

-- name: SelectUserAlreadyValidation :many
SELECT *
FROM users
where owner_validation is true
  and deleted is null;

-- name: UpdateUserOwnerValidation :one
update users
set owner_validation = true
where id_user = $1 RETURNING *;
;

-- name: UpdateLineIdByWhereUserID :one
UPDATE users SET id_line = $1 WHERE id_user = $2
RETURNING *;

