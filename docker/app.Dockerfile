FROM golang:1.22.2-alpine3.19

RUN apk update

WORKDIR /api

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest
