FROM golang:1.21.0

ENV CGO_ENABLED=0   \
    GOOS=linux      \
    GOARCH=amd64 

WORKDIR /usr/dev/go/delivery-service

COPY . .

RUN go mod download

RUN go build -o /delivery_app

CMD ["/delivery_app"]
