package main

import "context"

// Store struct that implements the OrdersStore interface
type Store struct {
	// mongoDB *mongo.Database
}

// NewStore creates a new Store instance
func NewStore() *Store {
	return &Store{
		// mongoDB: mongoDB,
	}
}

// Create method that implements the OrdersStore interface
func (s *Store) Create(context.Context) error {
	return nil
}
