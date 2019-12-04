package main

import (
	"fmt"

	customerpb "github.com/chriswilliams1977/moneytree/customer-service/proto/customer"
	"github.com/micro/go-micro"
)

func main() {

	repo := &Repository{}

	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.customer"),
	)

	srv.Init()

	h := &handler{repo}

	// Register our service with the gRPC server, this will tie our
	customerpb.RegisterCustomerServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
