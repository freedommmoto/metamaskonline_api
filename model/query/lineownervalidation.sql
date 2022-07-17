-- name: InsertLineOwnerValidation :one
INSERT INTO line_owner_validation (code)
VALUES ($1) RETURNING *;

-- name: SelectLastLineOwnerValidation :one
select *
from line_owner_validation
where id_user = $1
order by id_line_owner_validation
        desc limit 1;

-- name: SelectCodeUnConfirmWithIn3Houses :one
select *
from line_owner_validation
where true
  and created_at > now() - INTERVAL '180 minutes'
  and id_user is null
  and code = $1
order by id_line_owner_validation
        desc limit 1;

-- name: UpdateUserIDtoLineOwnerValidation :one
UPDATE line_owner_validation SET id_user = $1 WHERE id_line_owner_validation = $2 RETURNING *;
