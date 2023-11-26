# Curso Containers Completo & Total

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
