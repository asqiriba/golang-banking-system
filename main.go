package main

import (
	"database/sql"
	"log"

	"github.com/asqiriba/golang-banking-system/api"
	db "github.com/asqiriba/golang-banking-system/db/sqlc"
	"github.com/asqiriba/golang-banking-system/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Failed to open database ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start server ", err)
	}
}
