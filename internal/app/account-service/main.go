package main

import (
	"fmt"

	accountpb "github.com/chriswilliams1977/moneytree-protos/account"
	"github.com/micro/go-micro"
)


func main() {

	repo := &Repository{}

	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("moneytree.svc.account"),
	)

	srv.Init()



	// Register our implementation with 
	accountpb.RegisterAccountServiceHandler(srv.Server(), &handler{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
