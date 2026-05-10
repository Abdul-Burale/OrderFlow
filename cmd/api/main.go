package main

import (
	"log"

	"github.com/abdul-burale/orderflow/internal/config"
	"github.com/abdul-burale/orderflow/internal/server"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	/*
		dbConn, err := db.NewPostGres(cfg.DBConn)
		if err != nil {
			log.Fatal(err)
		}
	*/

	srv := server.New(cfg)
	server.Run(srv)
}
