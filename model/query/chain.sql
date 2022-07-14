-- name: selectChainByID :one
select *
from chain
where id_chain = $1
and deleted is null;

-- name: selectChainByChainCode :one
select *
from chain
where chain_code = $1
and deleted is null;