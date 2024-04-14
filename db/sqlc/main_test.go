package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver       = "pgx"
	dbSourceString = "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSourceString)
	if err != nil {
		log.Fatalf("cannot connect to db. err: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
