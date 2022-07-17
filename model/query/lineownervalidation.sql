-- name: InsertLineOwnerValidation :one
INSERT INTO line_owner_validation (code, id_user)
VALUES ($1, $2) RETURNING *;

-- name: SelectLineOwnerValidation :one
select *
from line_owner_validation
where id_line_owner_validation = $1
and deleted is null;

-- name: SelectCodeUnConfirmWithIn3Houses :one
select *
from line_owner_validation
where true
  and created_at > now() - INTERVAL '180 minutes'
  and deleted is null
  and code = $1
order by id_line_owner_validation
        desc limit 1;

-- name: UpdateUserIDtoLineOwnerValidation :one
UPDATE line_owner_validation SET id_user = $1 WHERE id_line_owner_validation = $2 RETURNING *;

-- name: DeleteLineOwnerValidation :one
UPDATE line_owner_validation SET deleted = now() WHERE id_line_owner_validation = $1 RETURNING *;
