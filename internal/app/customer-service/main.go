package main

import (
	"fmt"

	//make sure all related files are using same path to proto
	pb "github.com/chriswilliams1977/moneytree-protos/customer"
	"github.com/micro/go-micro"
)

func main() {

	repo := &Repository{}

	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("moneytree.svc.customer"),
	)

	srv.Init()

	h := &handler{repo}

	// Register our service with the gRPC server, this will tie our
	pb.RegisterCustomerServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
