# syntax=docker/dockerfile:1

#	pack Hello World Application image
FROM alpine:latest

WORKDIR /hello-world

COPY cmd/helloWorld.sh ./bin/

ENTRYPOINT ["/bin/sh", "-c", "/hello-world/bin/helloWorld.sh"]
