package main

import (
	"context"
	"encoding/json"

	"io/ioutil"
	"log"
	"os"

	customerpb "github.com/chriswilliams1977/moneytree-protos/customer"
	micro "github.com/micro/go-micro"
)

const (
	customerFilename = "customer.json"
	accountFilename  = "accounts.json"
)

//The customer client calls services to return data about the customer
func main() {

	//Set up reference to customer service
	customerClient := micro.NewService(micro.Name("go.micro.srv.customer.client"))
	customerClient.Init()

	customerService := customerpb.NewCustomerServiceClient("moneytree.svc.customer", customerClient.Client())

	//Initialize repos
	createCustomer(customerService)

	//Example service calls
	getCustomers(customerService)
	getCustomerById(customerService, "NL57ABNA00000000")
	getCustomerAccount(customerService, "NL57ABNA00000000")
}

func createCustomer(customerService customerpb.CustomerServiceClient) {

	// Contact the server and print out its response.
	file := customerFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	//Dummy data for a customer
	customer, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	//Create customer in memory
	createResponse, err := customerService.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create customer: %v", err)
	}
	log.Println("Customer created: ", createResponse.Customer.AccountNumber)
}


func getCustomers(customerService customerpb.CustomerServiceClient) {

	//Get customers from memory
	getAll, err := customerService.GetCustomers(context.Background(), &customerpb.Request{})
	if err != nil {
		log.Fatalf("Could not list customers: %v", err)
	}
	log.Println("Getting customers")
	for _, v := range getAll.Customers {
		log.Println(v)
	}
}

//getCustomerById ...
func getCustomerById(customerService customerpb.CustomerServiceClient, id string) {

	log.Println("Get customer account id: ", &customerpb.Request{AccountNumber:id})

	//Get customer from memory
	customer, err := customerService.GetCustomerById(context.Background(), &customerpb.Request{AccountNumber:id})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	log.Println("Customer found: ", customer)

}

func getCustomerAccount(customerService customerpb.CustomerServiceClient, id string){

	log.Println("Fetching Accounts ")

	//Get customer from memory
	account, err := customerService.GetCustomerAccounts(context.Background(), &customerpb.Request{AccountNumber:id})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	log.Println("Account found: ", account)
}

func parseFile(file string) (*customerpb.Customer, error) {
	var customer *customerpb.Customer
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &customer)
	return customer, err
}
