package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngenohkevin/go-auth/api"
	store "github.com/ngenohkevin/go-auth/db/store"
	"github.com/ngenohkevin/go-auth/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	stores := store.NewStore(connPool)

	server, err := api.NewServer(config, &stores)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
