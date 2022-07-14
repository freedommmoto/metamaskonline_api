// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: chain.sql

package model

import (
	"context"
)

const selectChainByChainCode = `-- name: SelectChainByChainCode :one
select id_chain, chain_code, chain_name, url_api, created_at, modified, deleted
from chain
where chain_code = $1
and deleted is null
`

func (q *Queries) SelectChainByChainCode(ctx context.Context, chainCode string) (Chain, error) {
	row := q.db.QueryRowContext(ctx, selectChainByChainCode, chainCode)
	var i Chain
	err := row.Scan(
		&i.IDChain,
		&i.ChainCode,
		&i.ChainName,
		&i.UrlApi,
		&i.CreatedAt,
		&i.Modified,
		&i.Deleted,
	)
	return i, err
}

const selectChainByID = `-- name: SelectChainByID :one
select id_chain, chain_code, chain_name, url_api, created_at, modified, deleted
from chain
where id_chain = $1
and deleted is null
`

func (q *Queries) SelectChainByID(ctx context.Context, idChain int32) (Chain, error) {
	row := q.db.QueryRowContext(ctx, selectChainByID, idChain)
	var i Chain
	err := row.Scan(
		&i.IDChain,
		&i.ChainCode,
		&i.ChainName,
		&i.UrlApi,
		&i.CreatedAt,
		&i.Modified,
		&i.Deleted,
	)
	return i, err
}