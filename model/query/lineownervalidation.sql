-- name: InsertLineOwnerValidation :one
INSERT INTO line_owner_validation (code, id_user)
VALUES ($1, $2) RETURNING *;

-- name: SelectLastLineOwnerValidation :one
select *
from line_owner_validation
where id_user = $1
order by id_line_owner_validation
        desc limit 1;
