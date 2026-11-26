package main

import (
	"context"
	"log"

	pb "github.com/BabarMaqsood10/commons/api"

	common "github.com/BabarMaqsood10/commons"
)

// Service struct that implements the OrdersService interface
// OrderStore is here because the Service struct needs to call the Create method from the Store struct, which implements the OrdersStore interface
type Service struct {
	store OrdersStore
}

// NewService creates a new Service instance with the provided OrdersStore
func NewService(store OrdersStore) *Service {
	return &Service{
		store: store,
	}
}

// CreateOrder method that implements the OrdersService interface
func (s *Service) CreateOrder(context.Context) error {
	return nil
}

func (s *Service) ValidateOrders(ctx context.Context, req *pb.CreateOrderRequest) error {
	if len(req.Items) == 0 {
		return common.ErrNoitems
	}
	mergedItems := mergeItemsQuantities(req.Items)

	// validate with stock service
	log.Printf("merged items: %v", mergedItems)
	return nil

}

func mergeItemsQuantities(items []*pb.ItemWithQuantity) []*pb.ItemWithQuantity {
	merged := make([]*pb.ItemWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.Id == item.Id {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			merged = append(merged, item)
		}
	}
	return merged
}
