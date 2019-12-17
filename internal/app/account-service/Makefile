# account-service/Makefile
#use paths=source_relative if you want generated proto relative --go_out=plugins=micro,paths=source_relative:. 
#list all go modules used go list -f "{{ .Path }} {{ .Dir }}" -m all 
build:
	docker build -t account .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 account