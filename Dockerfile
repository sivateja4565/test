
FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY *.go ./
COPY * ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /siva-test

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/siva-test"]
