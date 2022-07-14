package model

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/metamaskonline?sslmode=disable"
)

func TestMain(m *testing.M) {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println("unable to connect db:", err)
	}
	testQueries = New(connection)
	os.Exit(m.Run())
}
