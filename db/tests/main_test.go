package tests

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/ngenohkevin/go-auth/db/sqlc"
	"github.com/ngenohkevin/go-auth/utils"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = db.New(connPool)
	os.Exit(m.Run())
}
