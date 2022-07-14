-- name: InsertChainEvent :one
INSERT INTO chain_event ( wallet_id, action_type, from_metamask_wallet_id, to_metamask_wallet_id, log_event)
VALUES ( $1, $2, $3, $4, $5) RETURNING *;
