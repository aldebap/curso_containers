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

Para executar essa aplicação a partig a sua imagem, utilizar o seguinte comando:

```sh
docker run hello-world
```

### web-server

Essa pasta contém um aplicação web que utiliza o NginX, que é um servidor web, para publicar uma
única página HTML estática.

Para criar um container a partir dessa aplicação e publicá-lo como um pacote no GitHub, utilizar o
seguinte comando:

```sh
export  GITHUB_USER='aldebap'

docker build --tag ghcr.io/${GITHUB_USER}/hello-web --file Dockerfile .

export PAT='${GitHubs Personal Access Token}'
echo ${PAT} | docker login ghcr.io -u ${GITHUB_USER} --password-stdin
docker push ghcr.io/${GITHUB_USER}/hello-web

```

Se quiser utilizar o seu usuário de Github, deve-se substituir o valor da variável
```$GITHUB_USER``` no script acima, bem como criar um token PAT para o seu usuário,
utilizando esse token na variável ```$PAT```

Para executar essa aplicação a partir da sua imagem, utilizar o seguinte comando:

```sh
docker run -d -p 8080:80 ghcr.io/${GITHUB_USER}/hello-web
```

### rest-api

Essa pasta contém uma implementação de um endpoint rest escrita em Golang.
