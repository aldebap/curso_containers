# Curso Containers Completo & Total

[![Docker Publish](https://github.com/aldebap/curso_containers/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/aldebap/curso_containers/actions/workflows/docker-publish.yml)

Material de apoio para o Curso Containers Completo & Total da [CodeOptera](https://www.youtube.com/@CodeOptera)

### hello-world

Essa pasta contém uma simples aplicação Hello World que utiliza um Shell script (sh) para imprimir
a famosa mensagem "__Hello World__" na console.

Para criar um container a partir dessa aplicação, utilizar o seguinte comando:

```sh
docker build --tag hello-world --file Dockerfile .
```

Para criar um container a partir dessa aplicação e publicá-lo como um pacote no GitHub, utilizar o
seguinte comando:

```sh
docker build --tag ghcr.io/aldebap/hello-world --file Dockerfile .

export PAT='${GitHubs Personal Access Token}'
echo $PAT | docker login ghcr.io -u aldebap --password-stdin
docker push ghcr.io/aldebap/hello-world

```

Para executar essa aplicação a partig a sua imagem, utilizar o seguinte comando:

```sh
docker run hello-world
```

### web-server

Essa pasta contém um aplicação web que utiliza o NginX, que é um servidor web, para publicar uma
única página HTML estática.

Para criar um container a partir dessa aplicação, utilizar o seguinte comando:

```sh
docker build --tag hello-web --file Dockerfile .
```

Para criar um container a partir dessa aplicação e publicá-lo como um pacote no GitHub, utilizar o
seguinte comando:

```sh
docker build --tag ghcr.io/aldebap/hello-web --file Dockerfile .

export PAT='${GitHubs Personal Access Token}'
echo $PAT | docker login ghcr.io -u aldebap --password-stdin
docker push ghcr.io/aldebap/hello-web

```

Para executar essa aplicação a partig a sua imagem, utilizar o seguinte comando:

```sh
docker run -d -p 8080:80 hello-web
```
