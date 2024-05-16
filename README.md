# Instruções

## Executar com Docker Compose

Para iniciar a app e o redis com docker execute o seguinte comando:

```sh
docker-compose up -d
```

## Executar Testes Unitários da Aplicação

Para rodar os testes unitários da aplicação, use o comando:

```sh
go test ./...
```

## Executar o Projeto Localmente

Para executar o projeto localmente, utilize o comando:

```sh
go run ./main.go
```

## Executar o Projeto Dentro do Docker Manualmente (not recommend)

1. **Buildar a Imagem Docker**:

   Para construir a imagem Docker, execute:

   ```sh
   docker build -t rate-limiter:latest .
   ```

2. **Executar o Contêiner Docker:**:

   Para executar o contêiner Docker, mapeando a porta 8080 do host para a porta 8080 do contêiner, use:

   ```sh
   docker run -d -p 8080:8080 --name pos-go-redis-limiter rate-limiter:latest
   ```

