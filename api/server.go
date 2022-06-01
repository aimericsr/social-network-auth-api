package api

import (
	"fmt"
	"net/http"

	"github.com/aimericsr/social-network-auth-api/api/controller/generic"
	"github.com/aimericsr/social-network-auth-api/token"
	"github.com/aimericsr/social-network-auth-api/util"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var Serv *Server

type Server struct {
	Config     util.Config
	TokenMaker token.Maker
	DB         *gorm.DB
	Router     *mux.Router
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

	server := &Server{
		Config:     config,
		TokenMaker: tokenMaker,
		DB:         nil,
	}

	server.setupRouter()

	Serv = server

	fmt.Printf("Server listening on port %v\n", config.ServerAddress)
	err = http.ListenAndServe(":"+config.ServerAddress, server.Router)
	if err != nil {
		return fmt.Errorf("Cannot start api on the port:", err)
	}

	return nil
}

func (server *Server) setupRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/health", generic.HealthCheck).Methods("GET")

	// router.HandleFunc("/users", socialnetwork.CreateAccount).Methods("POST")

	// router.HandleFunc("/users/protected", middleware.BasicAuth(socialnetwork.CreateAccount, server.TokenMaker)).Methods("POST")

	server.Router = router
}
