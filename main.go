package main

import (
	"database/sql"
	"log"

	"github.com/henryeffiong/gobank/api"
	db "github.com/henryeffiong/gobank/db/sqlc"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver       = "pgx"
	dbSourceString = "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable"
	serverAddress  = "localhost:8080"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSourceString)
	if err != nil {
		log.Fatalf("cannot connect to db. err: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
