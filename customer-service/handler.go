package main

import (
	"context"

	pb "github.com/chriswilliams1977/moneytree/customer-service/proto/customer"
)

type handler struct {
	repository
}

// CreateCustomer - we created just one method on our service,
func (s *handler) CreateCustomer(ctx context.Context, req *pb.Customer, res *pb.Response) error {

	// Save our customer
	customer, err := s.repository.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Customer = customer
	return nil
}

// GetCustomers -
func (s *handler) GetCustomers(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	customers := s.repository.GetAll()
	res.Customers = customers
	return nil
}
