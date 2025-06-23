package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/john9101/simplebank/api"
	db "github.com/john9101/simplebank/db/sqlc"
	"github.com/john9101/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config:",err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
