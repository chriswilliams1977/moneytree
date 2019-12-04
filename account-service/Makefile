# account-service/Makefile
build:
	protoc -I. --go_out=plugins=micro:. \
      proto/account/account.proto
	docker build -t account .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 account