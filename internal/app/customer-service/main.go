package main

import (
	"fmt"

	//make sure all related files are using same path to proto
	accountpb "github.com/chriswilliams1977/moneytree-protos/account"
	customerpb "github.com/chriswilliams1977/moneytree-protos/customer"
	"github.com/micro/go-micro"
)

func main() {

	repo := &Repository{}

	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("moneytree.svc.customer"),
	)
	

	srv.Init()

	accountClient := accountpb.NewAccountServiceClient("moneytree.svc.account", srv.Client())

	h := &handler{repo, accountClient}

	// Register our service with the gRPC server, this will tie our
	customerpb.RegisterCustomerServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
