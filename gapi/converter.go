package gapi

import (
	db "github.com/henryeffiong/gobank/db/sqlc"
	"github.com/henryeffiong/gobank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func converter(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
