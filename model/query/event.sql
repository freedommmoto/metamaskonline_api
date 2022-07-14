-- name: InsertEvent :one
INSERT INTO event ( id_line_event, id_chain_event)
VALUES ( $1, $2) RETURNING *;