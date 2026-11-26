package main

import (
	"context"
	"log"

	pb "github.com/BabarMaqsood10/commons/api"
	"google.golang.org/grpc"
)

// grpcHandler struct that implements the OrderServiceServer interface
type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

// NEWGRPCHandler function that registers the OrderServiceServer with the provided gRPC server
func NEWGRPCHandler(grpcServer *grpc.Server) {
	// Create a new instance of grpcHandler
	handler := &grpcHandler{}
	// Register the OrderServiceServer with the provided gRPC server using the handler instance
	pb.RegisterOrderServiceServer(grpcServer, handler)

}

// CreateOrder method that implements the OrderServiceServer interface
// It receives a CreateOrderRequest and returns an Order response or an error
func (h *grpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	// Log a message indicating that an order has been received
	log.Printf("new order received. order %v", req)
	// Create a new Order instance with a hardcoded ID of "123"
	o := &pb.Order{
		Id: "123",
	}
	// Return the created Order instance and a nil error
	return o, nil

}
