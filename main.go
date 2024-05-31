package main

import (
	"database/sql"
	"log"
	"net"

	db "github.com/henryeffiong/gobank/db/sqlc"
	"github.com/henryeffiong/gobank/gapi"
	"github.com/henryeffiong/gobank/pb"
	"github.com/henryeffiong/gobank/util"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to load env: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db. err: %v", err)
	}

	store := db.NewStore(conn)
	// runGinServer(config, store)
	runGRPCServer(config, store)
}

// func runGinServer(config util.Config, store db.Store) {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatalf("cannot create server . err: %v", err)
// 	}

// 	err = server.Start(config.HTTPServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot start server: ", err)
// 	}
// }

func runGRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server . err: %v", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterGoBankServer(gRPCServer, server)
	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatalf("cannot create listener . err: %v", err)
	}

	log.Printf("starting gRPC server at %s...", listener.Addr().String())

	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server: ", err)
	}
}
