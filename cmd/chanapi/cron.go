package main

import (
	"database/sql"
	controller "github.com/freedommmoto/metamaskonline_api/controller"
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
	chText := make(chan string)
	if CainInfo.ChainCode == "bsc-testnet" {
		go controller.CallBSCCheckPerUser(chText, CronQueries, CainInfo)
	}
	if CainInfo.ChainCode == "eth-testnet" {
		go controller.CallETHCheckPerUser(chText, CronQueries, CainInfo)
	}
	chTextAll := <-chText // value from go Routine out with chanel data here
	log.Println("in function checkTransaction check wallet this list ", chTextAll)
}

var CronQueries *db.Queries
var CainInfo db.Chain

func main() {
	config, err := tool.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config file:", err)
		return
	}
	connection, errConnectDB := sql.Open(config.DBDriver, config.DBSource)
	if errConnectDB != nil {
		log.Println("can't connect database", errConnectDB)
	}
	CronQueries = db.New(connection)
	cain, err := lib.GetActiveChin(CronQueries)
	if err != nil {
		log.Println("GetActiveChin", err)
		return
	}
	CainInfo = cain
	doEverySetTime(time.Second*2, checkTransaction)
}
