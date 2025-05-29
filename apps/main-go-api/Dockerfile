# syntax=docker/dockerfile:1

ARG GO_VERSION=1.24.3

FROM golang:${GO_VERSION} AS builder

WORKDIR /usr/src/app

# Enable Go modules + no C dependencies
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux

COPY apps/main-go-api/go.mod apps/main-go-api/go.sum ./
RUN go mod download

COPY apps/main-go-api/*.go ./

RUN go build -o main-go-api


FROM builder AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=builder /usr/src/app/main-go-api ./main-go-api

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["./main-go-api"]