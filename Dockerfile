FROM golang:1.19.1-alpine3.15 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


# Build the application
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

EXPOSE 8080

RUN go build -o /main
CMD ["/main"]
