# Clean Architecture - Order System

Este projeto implementa um sistema de orders usando Clean Architecture em Go, com suporte a REST API, gRPC e GraphQL.

## Funcionalidades

- **Criar Order**: Endpoint para criar uma nova order
- **Listar Orders**: Endpoint para listar todas as orders
- **Três tipos de API**:
  - REST API (HTTP)
  - gRPC
  - GraphQL

## Pré-requisitos

- Go 1.19+
- Docker e Docker Compose
- Wire (para injeção de dependência)

## Instalação e Execução

### 1. Subir o banco de dados e RabbitMQ

```bash
docker compose up -d
```

### 2. Instalar dependências

```bash
go mod tidy
```

### 3. Gerar código Wire (injeção de dependência)

```bash
cd cmd/ordersystem
wire
```

### 4. Executar migrações do banco de dados

Execute o SQL em `sql/migrations/001_create_orders_table.sql` no banco MySQL.

### 5. Executar a aplicação

```bash
go run cmd/ordersystem/*.go
```

## Portas dos Serviços

- **REST API**: http://localhost:8000
- **gRPC Server**: localhost:50051
- **GraphQL**: http://localhost:8080
- **GraphQL Playground**: http://localhost:8080 (interface gráfica)

## Endpoints

### REST API

#### Criar Order
```http
POST http://localhost:8000/order
Content-Type: application/json

{
  "id": "order-001",
  "price": 100.0,
  "tax": 10.0
}
```

#### Listar Orders
```http
GET http://localhost:8000/order
```

### GraphQL

#### Criar Order
```graphql
mutation {
  createOrder(input: {
    id: "order-002"
    Price: 150.0
    Tax: 15.0
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

#### Listar Orders
```graphql
query {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC

Use um cliente gRPC para testar os serviços:

- `CreateOrder`: Cria uma nova order
- `ListOrders`: Lista todas as orders

## Estrutura do Projeto

```
CleanArch/
├── cmd/ordersystem/          # Comando principal da aplicação
├── internal/
│   ├── entity/               # Entidades de domínio
│   ├── usecase/              # Casos de uso
│   ├── infra/
│   │   ├── database/         # Repositórios
│   │   ├── web/              # Handlers REST
│   │   ├── grpc/             # Serviços gRPC
│   │   └── graph/            # Resolvers GraphQL
├── sql/migrations/           # Migrações do banco
├── api.http                  # Arquivo com requests de teste
└── docker-compose.yaml       # Configuração Docker
```

## Testando

Use o arquivo `api.http` para testar os endpoints REST e GraphQL diretamente no seu editor.

## Banco de Dados

O projeto usa MySQL como banco de dados principal. A estrutura da tabela `orders`:

- `id`: VARCHAR(255) PRIMARY KEY
- `price`: DECIMAL(10,2)
- `tax`: DECIMAL(10,2)
- `final_price`: DECIMAL(10,2)
- `created_at`: TIMESTAMP
