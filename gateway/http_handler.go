package main

import (
	"errors"
	"net/http"

	common "github.com/BabarMaqsood10/commons"
	pb "github.com/BabarMaqsood10/commons/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler struct that implements the HTTP handler for the gateway
type Handler struct {
	client pb.OrderServiceClient
}

// NewHandler creates a new Handler instance
func NewHandler(client pb.OrderServiceClient) *Handler {
	return &Handler{
		client: client,
	}
}

// registerRoutes method that registers the routes for the HTTP handler
func (h *Handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.HandleCreateOrder)

}

// HandleCreateOrder method that handles the HTTP request for creating an order
func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// Extract the customerId from the URL path
	customerID := r.PathValue("customerId")

	// Read the request body and unmarshal it into a slice of ItemWithQuantity
	var items []*pb.ItemWithQuantity
	if err := common.ReadJson(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Call the CreateOrder method on the gRPC client
	// Create a new CreateOrderRequest with the extracted customerId and items
	order, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerID,
		Items:      items,
	})
	// If there is an error, convert it to a gRPC status and handle it accordingly
	rStatus := status.Convert(err)
	// If the error is not nil, check the gRPC status code and write an appropriate HTTP response
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write the response as JSON with a status code of 201 Created
	common.WriteJson(w, http.StatusCreated, order)
}

func validateItems(items []*pb.ItemWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoitems
	}
	for _, item := range items {
		if item.Id == "" {
			return errors.New("itemId cannot be empty")
		}
		if item.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
	}
	return nil
}
