package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/zura-t/simplebank/api"
	db "github.com/zura-t/simplebank/db/sqlc"
	"github.com/zura-t/simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can't connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}
