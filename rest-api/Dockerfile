# syntax=docker/dockerfile:1

#	build the application
FROM golang:1.17-alpine AS build

WORKDIR /rest-api

COPY main.go go.mod go.sum ./

RUN CGO_ENABLED=0 go build -o ./bin/server main.go

#	create application image
FROM alpine:latest

WORKDIR /rest-api

COPY --from=build /rest-api/bin/server ./bin/

EXPOSE 8080
ENTRYPOINT ["/rest-api/bin/server"]
