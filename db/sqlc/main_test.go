package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/henryeffiong/gobank/util"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("unable to load env: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db. err: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
