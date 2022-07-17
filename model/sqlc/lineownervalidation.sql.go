// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: lineownervalidation.sql

package model

import (
	"context"
)

const deleteLineOwnerValidation = `-- name: DeleteLineOwnerValidation :one
UPDATE line_owner_validation SET deleted = now() WHERE id_line_owner_validation = $1 RETURNING id_line_owner_validation, code, id_user, created_at, deleted
`

func (q *Queries) DeleteLineOwnerValidation(ctx context.Context, idLineOwnerValidation int32) (LineOwnerValidation, error) {
	row := q.db.QueryRowContext(ctx, deleteLineOwnerValidation, idLineOwnerValidation)
	var i LineOwnerValidation
	err := row.Scan(
		&i.IDLineOwnerValidation,
		&i.Code,
		&i.IDUser,
		&i.CreatedAt,
		&i.Deleted,
	)
	return i, err
}

const insertLineOwnerValidation = `-- name: InsertLineOwnerValidation :one
INSERT INTO line_owner_validation (code, id_user)
VALUES ($1, $2) RETURNING id_line_owner_validation, code, id_user, created_at, deleted
`

type InsertLineOwnerValidationParams struct {
	Code   string `json:"code"`
	IDUser int32  `json:"id_user"`
}

func (q *Queries) InsertLineOwnerValidation(ctx context.Context, arg InsertLineOwnerValidationParams) (LineOwnerValidation, error) {
	row := q.db.QueryRowContext(ctx, insertLineOwnerValidation, arg.Code, arg.IDUser)
	var i LineOwnerValidation
	err := row.Scan(
		&i.IDLineOwnerValidation,
		&i.Code,
		&i.IDUser,
		&i.CreatedAt,
		&i.Deleted,
	)
	return i, err
}

const selectCodeUnConfirmWithIn3Houses = `-- name: SelectCodeUnConfirmWithIn3Houses :one
select id_line_owner_validation, code, id_user, created_at, deleted
from line_owner_validation
where true
  and created_at > now() - INTERVAL '180 minutes'
  and deleted is null
  and code = $1
order by id_line_owner_validation
        desc limit 1
`

func (q *Queries) SelectCodeUnConfirmWithIn3Houses(ctx context.Context, code string) (LineOwnerValidation, error) {
	row := q.db.QueryRowContext(ctx, selectCodeUnConfirmWithIn3Houses, code)
	var i LineOwnerValidation
	err := row.Scan(
		&i.IDLineOwnerValidation,
		&i.Code,
		&i.IDUser,
		&i.CreatedAt,
		&i.Deleted,
	)
	return i, err
}

const selectLineOwnerValidation = `-- name: SelectLineOwnerValidation :one
select id_line_owner_validation, code, id_user, created_at, deleted
from line_owner_validation
where id_line_owner_validation = $1
and deleted is null
`

func (q *Queries) SelectLineOwnerValidation(ctx context.Context, idLineOwnerValidation int32) (LineOwnerValidation, error) {
	row := q.db.QueryRowContext(ctx, selectLineOwnerValidation, idLineOwnerValidation)
	var i LineOwnerValidation
	err := row.Scan(
		&i.IDLineOwnerValidation,
		&i.Code,
		&i.IDUser,
		&i.CreatedAt,
		&i.Deleted,
	)
	return i, err
}

const updateUserIDtoLineOwnerValidation = `-- name: UpdateUserIDtoLineOwnerValidation :one
UPDATE line_owner_validation SET id_user = $1 WHERE id_line_owner_validation = $2 RETURNING id_line_owner_validation, code, id_user, created_at, deleted
`

type UpdateUserIDtoLineOwnerValidationParams struct {
	IDUser                int32 `json:"id_user"`
	IDLineOwnerValidation int32 `json:"id_line_owner_validation"`
}

func (q *Queries) UpdateUserIDtoLineOwnerValidation(ctx context.Context, arg UpdateUserIDtoLineOwnerValidationParams) (LineOwnerValidation, error) {
	row := q.db.QueryRowContext(ctx, updateUserIDtoLineOwnerValidation, arg.IDUser, arg.IDLineOwnerValidation)
	var i LineOwnerValidation
	err := row.Scan(
		&i.IDLineOwnerValidation,
		&i.Code,
		&i.IDUser,
		&i.CreatedAt,
		&i.Deleted,
	)
	return i, err
}
