package main

import (
	"context"
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/lib"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	tool "github.com/freedommmoto/metamaskonline_api/tool"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func doEverySetTime(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func checkTransaction(t time.Time) {
	ch := make(chan string)
	go callCheckPerUser(ch)
	messageFromMisterB := <-ch // value from go Routine out with chanel data here
	log.Println("in function checkTransaction check wallet this list ", messageFromMisterB)
}

func callCheckPerUser(message chan<- string) {
	userList, err := CronQueries.SelectUserAlreadyValidation(context.Background())
	if err != nil {
		log.Println("error SelectUserAlreadyValidation", err)
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
			log.Println("error SelectFollowWalletByIDUser", err)
			return
		}
		walletALL += wallet.MetamaskWalletID + ","
	}
	//log.Println(wallet.MetamaskWalletID)
	message <- walletALL
}

var CronQueries *db.Queries

func main() {
	config, err := tool.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config file:", err)
		return
	}
	lib.DefaultBNBPrice = 250.0
	connection, errConnectDB := sql.Open(config.DBDriver, config.DBSource)
	if errConnectDB != nil {
		log.Println("can't connect database", errConnectDB)
	}
	CronQueries = db.New(connection)

	doEverySetTime(time.Second, checkTransaction)
}
