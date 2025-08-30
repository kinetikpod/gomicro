package main

import (
	"log"

	"github.com/kinetikpod/gomicro/order/config"
	"github.com/kinetikpod/gomicro/order/internal/adapters/db"
	"github.com/kinetikpod/gomicro/order/internal/adapters/grpc"
	"github.com/kinetikpod/gomicro/order/internal/application/core/api"
)

func main() {
	// Init DB adapter pakai DATA_SOURCE_URL dari env
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	// Core application (business logic)
	application := api.NewApplication(dbAdapter)

	// gRPC adapter, pakai APPLICATION_PORT dari env
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	// Run gRPC service
	grpcAdapter.Run()
}
