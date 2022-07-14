-- name: SelectChainByID :one
select *
from chain
where id_chain = $1
and deleted is null;

-- name: SelectChainByChainCode :one
select *
from chain
where chain_code = $1
and deleted is null;