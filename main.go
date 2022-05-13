package main

import (
	"log"

	"github.com/aimericsr/social-network-auth-api/util"
)

var server *server.

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config variables:", err)
	}
	serv, err := server.NewServer(config)
	if err != nil {
		log.Fatal("Can't start the server:", err)
	}
}
