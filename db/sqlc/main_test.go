package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQueries *Queries

const (
	dbDriver       = "pgx"
	dbSourceString = "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSourceString)
	if err != nil {
		log.Fatalf("cannot connect to db. err: %v", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
