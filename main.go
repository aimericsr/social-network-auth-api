package main

import (
	"log"

	"github.com/aimericsr/social-network-auth-api/api"
	"github.com/aimericsr/social-network-auth-api/db"
	"github.com/aimericsr/social-network-auth-api/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config variables:", err)
	}

	db.MakeMigration(config.DBSource)

	err = api.NewServer(config)
	if err != nil {
		log.Fatal("Can't start the server:", err)
	}
}
