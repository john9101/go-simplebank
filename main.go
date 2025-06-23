package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/john9101/simplebank/api"
	db "github.com/john9101/simplebank/db/sqlc"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)
	
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
