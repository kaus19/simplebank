package main

import (
	"database/sql"
	"log"

	"github.com/kaus19/simplebank/api"
	db "github.com/kaus19/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = ":8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
