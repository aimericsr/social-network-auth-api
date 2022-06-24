package server

import (
	"fmt"
	"net/http"

	"github.com/aimericsr/social-network-auth-api/api/handler"
	"github.com/aimericsr/social-network-auth-api/db"
	"github.com/aimericsr/social-network-auth-api/token"
	"github.com/aimericsr/social-network-auth-api/util"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

var Serv *Server

type Server struct {
	Config     util.Config
	TokenMaker token.Maker
	DB         *gorm.DB
	Router     *chi.Mux
}

func NewServer(config util.Config) (err error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker: %w", err)
	}

	// DB, err := db.Open(config.DBSource)
	// if err != nil {
	// 	return fmt.Errorf("Cannot connect to database:", err)
	// }

	db.MakeMigration(config.DBSource)

	router := handler.NewHandler(tokenMaker)

	server := &Server{
		Config:     config,
		TokenMaker: tokenMaker,
		DB:         nil,
		Router:     router,
	}

	Serv = server

	err = http.ListenAndServe(":"+config.ServerAddress, server.Router)
	if err != nil {
		return err
	}
	fmt.Printf("Server listening on port %v\n", config.ServerAddress)
	return nil
}
