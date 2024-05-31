package gapi

import (
	db "github.com/henryeffiong/gobank/db/sqlc"
	"github.com/henryeffiong/gobank/pb"
	"github.com/henryeffiong/gobank/token"
	"github.com/henryeffiong/gobank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedGoBankServer // All implementations must embed UnimplementedGoBankServer for forward compatibility
	config                       util.Config
	store                        db.Store
	tokenMaker                   token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
