#To call server
#Add import: micro "github.com/micro/go-micro"
#Declare the service: service := micro.NewService(micro.Name("go.micro.srv.customer.client"))
#Initialize service: service.Init()
#Create new client: client := pb.NewCustomerServiceClient("go.micro.srv.customer", service.Client())
#Note: Import protobufs from external repo
build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t customer-client .

run:
	docker run customer-client