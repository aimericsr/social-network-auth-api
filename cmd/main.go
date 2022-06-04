package main

import (
	"hexa-archi/internal/adapters/app/api"
	"hexa-archi/internal/adapters/core/arithmetic"
	gRPC "hexa-archi/internal/adapters/framework/left/grpc"
	"hexa-archi/internal/adapters/framework/right/db"
	"hexa-archi/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	//ports
	var dbasAdapter ports.DbPort
	var core ports.ArithmetiquePort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbasAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate database connection: %v", err)
	}
	defer dbasAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbasAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
