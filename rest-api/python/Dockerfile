# syntax=docker/dockerfile:1

#	build the application
FROM microsoft/nanoserver AS build

WORKDIR /rest-api

COPY main.py service.py ./

RUN CGO_ENABLED=0 pip3 install Flask
RUN CGO_ENABLED=0 go build -o ./bin/server main.go

#	create application image
FROM microsoft/nanoserver

#   properties
LABEL org.opencontainers.image.source https://github.com/aldebap/rest-api-python

WORKDIR /rest-api

COPY --from=build /rest-api/bin/server ./bin/

EXPOSE 8080
ENTRYPOINT ["/rest-api/bin/server"]
