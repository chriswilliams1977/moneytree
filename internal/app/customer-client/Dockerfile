FROM golang:1.13.0 as builder

ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build  -o moneytree-client-customer -a -installsuffix cgo client.go

FROM scratch

COPY --from=builder /app .

CMD ["./customer-client"]