package main

import (
	"log"

	"github.com/aimericsr/social-network-auth-api/api"
	"github.com/aimericsr/social-network-auth-api/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config variables:", err)
	}

	err = api.NewServer(config)
	if err != nil {
		log.Fatal("Can't start the server:", err)
	}
}
