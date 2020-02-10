FROM golang:1.13.0 as builder

ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  -o moneytree-service-customer -a -installsuffix cgo main.go repository.go handler.go

FROM scratch

COPY --from=builder /app .

CMD ["./account-service"]