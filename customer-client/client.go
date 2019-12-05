package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/chriswilliams1977/protobufs/protos/customer"
	micro "github.com/micro/go-micro"
)

const (
	defaultFilename = "customer.json"
)

func parseFile(file string) (*pb.Customer, error) {
	var customer *pb.Customer
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &customer)
	return customer, err
}

func main() {

	service := micro.NewService(micro.Name("go.micro.srv.customer.client"))
	service.Init()

	//this is the service connecting too so name must be same as micro.Name("go.micro.srv.consignment")
	client := pb.NewCustomerServiceClient("go.micro.srv.customer", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	customer, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetCustomers(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Customers {
		log.Println(v)
	}
}
