package main

import (
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/controller"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	"github.com/freedommmoto/metamaskonline_api/tool"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq" //need this for connect postgres
	"log"
	"net/http"
	"time"
)

var ApiQueries *db.Queries

func main() {
	//set time-zone
	loc, errTime := time.LoadLocation("UTC")
	if errTime != nil {
		log.Println("unable to set time zone:", errTime)
	}
	time.Local = loc

	config, err := tool.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config file:", err)
		return
	}

	connection, errConnectDB := sql.Open(config.DBDriver, config.DBSource)
	if errConnectDB != nil {
		log.Println("can't connect database", errConnectDB)
	}
	ApiQueries = db.New(connection)

	e := echo.New()
	e.POST("/line-webhook", func(c echo.Context) error {
		err := controller.ReplyMessageLine(c, ApiQueries, config)
		if err != nil {
			log.Println("replyMessageLine error :", err)
			return c.String(http.StatusInternalServerError, "error")
		}
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":8888"))
}
