package server

import (
	"fmt"
	"log"

	"github.com/aimericsr/social-network-auth-api/api/middlware"
	"github.com/aimericsr/social-network-auth-api/db"
	"github.com/aimericsr/social-network-auth-api/token"
	"github.com/aimericsr/social-network-auth-api/util"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/aimericsr/social-network-auth-api/api/generic"
	socialNetwork "github.com/aimericsr/social-network-auth-api/api/social-network"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
	DB         *gorm.DB
	router     *mux.Router
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	DB, err := db.Open(config.DBSource)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to database:", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		DB:         DB,
	}

	server.setupRouter()

	fmt.Printf("Server listening on port %v\n", config.ServerAddress)
	err = http.ListenAndServe(":"+config.ServerAddress, router)
	if err != nil {
		return nil, fmt.Errorf("Cannot start api on the port:", err)
	}

	return server, nil
}

func (server *Server) setupRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/health", generic.HealthCheck).Methods("GET")

	router.HandleFunc("/users", socialNetwork.CreateAccount).Methods("POST")
	router.HandleFunc("/users/protected", middlware.BasicAuth(socialNetwork.CreateAccount, server.tokenMaker)).Methods("POST")

	server.router = router
}

f


