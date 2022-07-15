-- name: SelectWalletByID :one
select *
from wallet
where wallet_id = $1
and deleted is null;

-- name: SelectWalletByIDUser :many
select *
from wallet
where id_user = $1
and deleted is null;

-- name: SelectWalletByMetamaskWalletID :one
select *
from wallet
where metamask_wallet_id = $1
and deleted is null;

-- name: SelectFollowWalletByIDUser :one
select *
from wallet
where id_user = $1
and follow_wallet = true
and deleted is null;

-- name: UpdateLastBlockNumber :one
update wallet
set last_block_number = $1
where wallet_id = $2
and deleted is null
RETURNING *;
;


