package main

import (
	"context"
	"log"
	"net"

	common "github.com/BabarMaqsood10/commons"
	"google.golang.org/grpc"
)

// main function that starts the gRPC server and registers the services
var (
	// grpcAddr is the address that the gRPC server will listen on, defaulting to "localhost:2000"
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()
	// Register the OrderServiceServer with the gRPC server
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	// store from store.go, used for database operations,
	store := NewStore()
	// service from service.go, used for business logic,
	service := NewService(store)
	// Register the gRPC handler with the gRPC server
	NEWGRPCHandler(grpcServer)
	// call the CreateOrder method from service.go, which will eventually call the Create method from store.go
	service.CreateOrder(context.Background())

	log.Println("Starting gRPC server on", grpcAddr)
	// Start the gRPC server and log any errors that occur

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}

}
