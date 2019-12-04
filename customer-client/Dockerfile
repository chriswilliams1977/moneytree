FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD customer.json /app/customer.json
ADD customer-client /app/customer-client

CMD ["./customer-client"]