package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/chriswilliams1977/moneytree/customer-service/proto/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Default port for GRPC
const (
	port = ":50051"
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

	//reader/writer mutual exclusion lock
	// If a goroutine holds a RWMutex for reading and another goroutine might call Lock,
	//no goroutine should expect to be able to acquire a read lock until the initial read lock is released.
	//In particular, this prohibits recursive read locking.
	mu        sync.RWMutex
	customers []*pb.Customer
}

// Create a new consignment
func (repo *Repository) Create(customer *pb.Customer) (*pb.Customer, error) {

	//create a lock - no other goroutine can read until unlocked
	repo.mu.Lock()
	updated := append(repo.customers, customer)
	repo.customers = updated
	repo.mu.Unlock()
	return customer, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Customer {
	return repo.customers
}

// Service should implement all of the methods to satisfy the service
type service struct {
	repo repository
}

// CreateCustomer - we created just one method on our service,
func (s *service) CreateCustomer(ctx context.Context, req *pb.Customer) (*pb.Response, error) {

	// Save our customer
	customer, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{
		Created:  true,
		Customer: customer,
	}, nil
}

// GetCustomers -
func (s *service) GetCustomers(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	customers := s.repo.GetAll()
	return &pb.Response{Customers: customers}, nil
}

func main() {

	repo := &Repository{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	pb.RegisterCustomerServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	// gRPC Server Reflection provides information about publicly-accessible gRPC services on a server,
	// and assists clients at runtime to construct RPC requests and responses without precompiled service information.
	//It is used by gRPC CLI, which can be used to introspect server protos and send/receive test RPCs.
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
