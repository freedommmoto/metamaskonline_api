-- name: SelectWalletByID :one
select *
from wallet
where wallet_id = $1
and deleted is null;

-- name: SelectWalletByIDUser :many
select *
from wallet
where wallet_id = $1
and deleted is null;

-- name: SelectWalletByMetamaskWalletID :one
select *
from wallet
where wallet_id = $1
and deleted is null;

-- name: UpdateLastBlockNumber :one
update wallet
set last_block_number = $1
where wallet_id = $2
and deleted is null
RETURNING *;
;


