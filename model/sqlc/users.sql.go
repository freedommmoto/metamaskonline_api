// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: users.sql

package model

import (
	"context"
	"database/sql"
)

const selectUserByLineUserID = `-- name: SelectUserByLineUserID :one
select id_user, username, password, id_line, owner_validation, created_at, modified, deleted
from users
where id_line = $1
  and deleted is null
`

func (q *Queries) SelectUserByLineUserID(ctx context.Context, idLine sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, selectUserByLineUserID, idLine)
	var i User
	err := row.Scan(
		&i.IDUser,
		&i.Username,
		&i.Password,
		&i.IDLine,
		&i.OwnerValidation,
		&i.CreatedAt,
		&i.Modified,
		&i.Deleted,
	)
	return i, err
}

const selectUserID = `-- name: SelectUserID :one
select id_user, username, password, id_line, owner_validation, created_at, modified, deleted
from users
where id_user = $1
and deleted is null
`

func (q *Queries) SelectUserID(ctx context.Context, idUser int32) (User, error) {
	row := q.db.QueryRowContext(ctx, selectUserID, idUser)
	var i User
	err := row.Scan(
		&i.IDUser,
		&i.Username,
		&i.Password,
		&i.IDLine,
		&i.OwnerValidation,
		&i.CreatedAt,
		&i.Modified,
		&i.Deleted,
	)
	return i, err
}

const updateUserOwnerValidation = `-- name: UpdateUserOwnerValidation :one
update users
set owner_validation = true
where id_user = $1
RETURNING id_user, username, password, id_line, owner_validation, created_at, modified, deleted
`

func (q *Queries) UpdateUserOwnerValidation(ctx context.Context, idUser int32) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserOwnerValidation, idUser)
	var i User
	err := row.Scan(
		&i.IDUser,
		&i.Username,
		&i.Password,
		&i.IDLine,
		&i.OwnerValidation,
		&i.CreatedAt,
		&i.Modified,
		&i.Deleted,
	)
	return i, err
}
