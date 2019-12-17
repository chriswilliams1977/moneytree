FROM golang:1.13.0 as builder

ENV GO111MODULE=on

#TODO install git and clone proto repo git clone https://github.com/chriswilliams1977/moneytree-protos.git protos
#TODO then execute Makefile

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o account

FROM scratch

COPY --from=builder /app .

CMD ["./account"]