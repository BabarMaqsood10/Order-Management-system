package main

import (
	"context"

	pb "github.com/BabarMaqsood10/commons/api"
)

// OrdersService interface that defines the methods for the service layer
type OrdersService interface {
	CreateOrder(context.Context) error
	ValidateOrders(context.Context, *pb.CreateOrderRequest) error
}

// OrdersStore interface that defines the methods for the store layer
type OrdersStore interface {
	Create(context.Context) error
}
