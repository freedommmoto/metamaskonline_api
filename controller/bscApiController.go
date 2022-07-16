package controller

import (
	"context"
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/lib"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	"github.com/freedommmoto/metamaskonline_api/tool"
)

func CallETHCheckPerUser(message chan<- string, CronQueries *db.Queries, CainInfo db.Chain) {
	//not implement in this version but similar with bsc
}

func CallBSCCheckPerUser(message chan<- string, CronQueries *db.Queries, CainInfo db.Chain) {
	lib.DefaultBNBPrice = 250.0
	userList, err := CronQueries.SelectUserAlreadyValidation(context.Background())
	if err != nil {
		tool.AddErrorLogIntoFile(err.Error())
		return
	}
	walletALL := ""
	for i := range userList {
		user := userList[i]

		wallet, err := CronQueries.SelectFollowWalletByIDUser(context.Background(), user.IDUser)
		if err == sql.ErrNoRows {
			continue
		}
		if err != nil {
			tool.AddErrorLogIntoFile(err.Error())
			continue
		}

		walletALL += wallet.MetamaskWalletID + ","
	}
	message <- walletALL
}
