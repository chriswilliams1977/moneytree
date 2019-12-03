#Steps to integrate micro
#Update protoc build to use micro plugin e.g protoc -I. --go_out=plugins=micro:.
#Update service to use micro steps:
	#Add import "github.com/micro/go-micro"
	#Update service methods to take a request and response
	#In main add new micro service and use name from proto package :
	#srv := micro.NewService ( micro.Name("go.micro.srv.consignment"), )
#Initialize service : srv.Init()
#Register service: pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
#Run service: if err := srv.Run(); err != nil {
build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/customer/customer.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t customer .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 customer