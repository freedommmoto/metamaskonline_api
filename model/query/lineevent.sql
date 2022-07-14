-- name: InsertLineEvent :one
INSERT INTO line_event ( id_line_user, id_use, request_log_event, response_log_event, error, error_text)
VALUES ( $1, $2, $3, $4, $5, $6) RETURNING *;
