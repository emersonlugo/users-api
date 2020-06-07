FROM golang:1.14-alpine

MAINTAINER 'Emerson Lugo <emersonlugo@gmail.com>'

# Set the Current Working Directory inside the container
WORKDIR /go/src/users-api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./cmd/users-api .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./cmd/users-api"]


