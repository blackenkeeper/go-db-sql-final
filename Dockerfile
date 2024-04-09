FROM golang:1.21.0
WORKDIR /usr/dev/go/delivery-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /delivery_app
CMD ["/delivery_app"]
