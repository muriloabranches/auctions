# Auctions API

Este projeto é uma API para gerenciar leilões e lances. A API é construída em Go e utiliza MongoDB como banco de dados.
Baseado no projeto [labs-auction-goexpert](https://github.com/devfullcycle/labs-auction-goexpert)

## Pré-requisitos

- Docker
- Docker Compose

## Configuração

1. Clone o repositório:

```sh
git clone https://github.com/seu-usuario/auctions.git
cd auctions
```

2. Crie um arquivo `.env` no diretório `cmd/auction` com o seguinte conteúdo:

```
BATCH_INSERT_INTERVAL=20s
MAX_BATCH_SIZE=4
AUCTION_INTERVAL=20s

MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=admin
MONGODB_URL=mongodb://admin:admin@mongodb:27017/auctions?authSource=admin
MONGODB_DB=auctions
```

## Execução do Projeto

Para iniciar o projeto, execute o seguinte comando:

```sh
docker-compose up --build
```

A API estará disponível em `http://localhost:8080`.

## Execução dos Testes

Para executar os testes, utilize o seguinte comando:

```sh
go test ./...
```

## Exemplos de cURL

### Criar um novo leilão

```sh
curl -X POST http://localhost:8080/auction \
-H "Content-Type: application/json" \
-d '{
    "product_name": "Product 1",
    "category": "Category 1",
    "description": "Description 1",
    "condition": 1
}'
```

### Criar um novo lance para o leilão mais recente

```sh
curl -X POST http://localhost:8080/bid \
-H "Content-Type: application/json" \
-d '{
    "user_id": "00000000-0000-0000-0000-000000000000",
    "auction_id": "AUCTION_ID",
    "amount": 25.0
}'
```

### Obter o lance vencedor do leilão mais recente

```sh
curl -X GET http://localhost:8080/auction/winner/AUCTION_ID
```

### Obter informações do leilão mais recente

```sh
curl -X GET http://localhost:8080/auction/AUCTION_ID
```

### Listar todos os leilões

```sh
curl -X GET "http://localhost:8080/auction?status=0"
```

Substitua `AUCTION_ID` pelo ID do leilão que você deseja consultar.