package gapi

import (
	"fmt"

	db "github.com/john9101/go-simplebank/db/sqlc"
	"github.com/john9101/go-simplebank/pb"
	"github.com/john9101/go-simplebank/token"
	"github.com/john9101/go-simplebank/util"
)

type Server struct {
	pb.UnimplementedGoSimpleBankServer
	config      util.Config
	store       db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("can not create token marker: %d", err)
	}

	server := &Server{
		config:      config,
		store:       store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}