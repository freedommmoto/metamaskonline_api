package lib

import (
	"context"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	model "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	_ "github.com/lib/pq"
)

func GetActiveChin(CronQueries *db.Queries) (chain model.Chain, err error) {
	//for this version we support only bsc test net
	activeChin := 1
	chain, err = CronQueries.SelectChainByID(context.Background(), int32(activeChin))
	if err != nil {
		return
	}
	return
}
