package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aimericsr/social-network-auth-api/db"
	"github.com/aimericsr/social-network-auth-api/util"
	"github.com/gorilla/mux"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config variables:", err)
	}
	err = db.MakeMigration(config.DBSource)
	if err != nil {
		log.Fatal("Cannot make migration:", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/v1/health", healthCheck).Methods("GET")
	http.ListenAndServe(config.ServerAddress, router)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is alive")
}
