FROM golang:1.20.3-alpine3.17
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install github.com/go-delve/delve/cmd/dlv@latest
CMD ["go", "run", "server.go"]