// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: lineevent.sql

package model

import (
	"context"
	"database/sql"
)

const insertLineEvent = `-- name: InsertLineEvent :one
INSERT INTO line_event ( id_line_user, id_use, request_log_event, response_log_event, error, error_text)
VALUES ( $1, $2, $3, $4, $5, $6) RETURNING id_line_event, id_line_user, id_use, request_log_event, response_log_event, error, error_text, created_at
`

type InsertLineEventParams struct {
	IDLineUser       sql.NullString `json:"id_line_user"`
	IDUse            int32          `json:"id_use"`
	RequestLogEvent  sql.NullString `json:"request_log_event"`
	ResponseLogEvent sql.NullString `json:"response_log_event"`
	Error            sql.NullBool   `json:"error"`
	ErrorText        sql.NullString `json:"error_text"`
}

func (q *Queries) InsertLineEvent(ctx context.Context, arg InsertLineEventParams) (LineEvent, error) {
	row := q.db.QueryRowContext(ctx, insertLineEvent,
		arg.IDLineUser,
		arg.IDUse,
		arg.RequestLogEvent,
		arg.ResponseLogEvent,
		arg.Error,
		arg.ErrorText,
	)
	var i LineEvent
	err := row.Scan(
		&i.IDLineEvent,
		&i.IDLineUser,
		&i.IDUse,
		&i.RequestLogEvent,
		&i.ResponseLogEvent,
		&i.Error,
		&i.ErrorText,
		&i.CreatedAt,
	)
	return i, err
}
