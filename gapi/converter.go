package gapi

import (
	db "github.com/john9101/go-simplebank/db/sqlc"
	"github.com/john9101/go-simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func converterUser(user db.User) *pb.User{
	return &pb.User{
		Username: user.Username,
		Email: user.Email,
		FullName: user.FullName,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}