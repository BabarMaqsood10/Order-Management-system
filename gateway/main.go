package main

import (
	"log"
	"net/http"

	common "github.com/BabarMaqsood10/commons"
	pb "github.com/BabarMaqsood10/commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// httpAddr is the address that the HTTP server will listen on, defaulting to ":8080"
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")

	orderServiceAddr = common.EnvString("ORDER_SERVICE_ADDR", "localhost:2000")
)

// main function that starts the HTTP server and registers the routes
func main() {
	// grpc.Dial creates a client connection to the gRPC server at ORDER_SERVICE_ADDR.
	// This connection is used to send RPC calls to the order service.
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to order service:", err)
	}
	// defer ensures the gRPC connection is closed when the program exits.
	defer conn.Close()

	log.Println("Connected to order service at", orderServiceAddr)
	// Create a new OrderServiceClient using the gRPC connection. This client will be used to make RPC calls to the order service.
	c := pb.NewOrderServiceClient(conn)
	// create a new ServeMux to register the routes
	mux := http.NewServeMux()
	// create a new Handler instance and register the routes
	handler := NewHandler(c)
	// register the routes for the HTTP handler
	handler.registerRoutes(mux)
	// start the HTTP server on port 8080 and log any errors that occur

	log.Println("Starting server on ", httpAddr)
	// start the HTTP server on port 8080 and log any errors that occur
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start the server:", err)
	}
}
