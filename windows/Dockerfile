# syntax=docker/dockerfile:1

#	pack Windows Hello World Application image
FROM microsoft/nanoserver

WORKDIR /hello-world

COPY bin/helloWorld.ps1 ./bin/

ENTRYPOINT ["PowerShell", "/hello-world/bin/helloWorld.ps1"]
