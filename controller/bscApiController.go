package controller

import (
	"context"
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/lib"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	"github.com/freedommmoto/metamaskonline_api/tool"
	_ "github.com/lib/pq"
)

func CallETHCheckPerUser(message chan<- string, CronQueries *db.Queries, CainInfo db.Chain) {
	//not implement in this version but similar with bsc
}

func CallBSCCheckPerUser(
	message chan<- string,
	CronQueries *db.Queries,
	CainInfo db.Chain,
	BSCToken string,
	LineToken string,
) {
	lib.DefaultBNBPrice = 250.0
	userList, err := CronQueries.SelectUserAlreadyValidation(context.Background())
	if err != nil {
		tool.AddErrorLogIntoFile(err.Error())
		message <- err.Error()
		return
	}

	//get bnb best price
	bnbPrice, err := lib.GetLastPriceBNB(CainInfo)
	if err != nil {
		tool.AddErrorLogIntoFile(err.Error())
		message <- err.Error()
		return
	} else {
		lib.DefaultBNBPrice = bnbPrice
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

		resultTx, LastBlockFromApi, errCallBscScan := lib.GetLastBlockTransactionFromBscScan(CainInfo, wallet.MetamaskWalletID, wallet.LastBlockNumber, BSCToken)
		if errCallBscScan != nil {
			tool.AddErrorLogIntoFile("GetLastBlockTransactionFromBscScan :" + errCallBscScan.Error())
			continue
		}

		//if the last block id is more then old one make alert
		if LastBlockFromApi > wallet.LastBlockNumber {
			arg := db.UpdateLastBlockNumberParams{
				LastBlockNumber: LastBlockFromApi,
				WalletID:        wallet.WalletID,
			}

			//update last block into wallet table
			_, errUpdateLastBlockNumber := CronQueries.UpdateLastBlockNumber(context.Background(), arg)
			if errUpdateLastBlockNumber != nil {
				tool.AddErrorLogIntoFile("errUpdateLastBlockNumber :" + err.Error())
				continue
			}

			//send alert to line
			lineAlertStr, err := lib.GetPushTextForLineAlert(resultTx, wallet.MetamaskWalletID, wallet.WalletName.String)
			if err != nil {
				tool.AddErrorLogIntoFile("MakeReplyTextForLineAlert :" + err.Error())
				continue
			}
			errAfterCallLine := lib.MakePushOneUserWithLineAPI(lineAlertStr, user, LineToken)
			if errAfterCallLine != nil {
				tool.AddErrorLogIntoFile("MakePushOneUserWithLineAPI :" + errAfterCallLine.Error())
				continue
			}
		}

		walletALL += wallet.MetamaskWalletID + ","
	}
	message <- walletALL
}
