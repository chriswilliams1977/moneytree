package main

import (
	// Import the generated protobuf code
	pb "github.com/chriswilliams1977/moneytree-protos/customer"
)

//interface to define data a repo can handle
type repository interface {

	//takes a customer a returns and customer
	Create(*pb.Customer) (*pb.Customer, error)
	//return slice of customers
	GetAll() []*pb.Customer
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
//Create repo object and implement interface to match type
type Repository struct {
	customers []*pb.Customer
}

// Create a new consignment
func (repo *Repository) Create(customer *pb.Customer) (*pb.Customer, error) {
	updated := append(repo.customers, customer)
	repo.customers = updated
	return customer, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Customer {
	return repo.customers
}
