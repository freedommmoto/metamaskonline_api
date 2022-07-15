package main

import (
	"database/sql"
	"github.com/freedommmoto/metamaskonline_api/controller"
	db "github.com/freedommmoto/metamaskonline_api/model/sqlc"
	tool "github.com/freedommmoto/metamaskonline_api/tool"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq" //need this for connect postgres
	"log"
	"net/http"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/metamaskonline?sslmode=disable"
)

var ApiQueries *db.Queries

func main() {

	connection, errConnectDB := sql.Open(dbDriver, dbSource)
	if errConnectDB != nil {
		log.Println("can't connect database", errConnectDB)
	}
	ApiQueries = db.New(connection)

	e := echo.New()
	e.POST("/line-webhook", func(c echo.Context) error {
		err := controller.ReplyMessageLine(c, ApiQueries)
		if err != nil {
			log.Println("replyMessageLine error :", err)
			return c.String(http.StatusInternalServerError, "error")
		}
		str := tool.RandomString(22)
		log.Println("in line-webhook", str)
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":8888"))
}
