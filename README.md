do# Clean Architecture - Order System

Este projeto implementa um sistema de orders usando Clean Architecture em Go, com suporte a REST API, gRPC e GraphQL.

## Funcionalidades

- **Criar Order**: Endpoint para criar uma nova order
- **Listar Orders**: Endpoint para listar todas as orders (principal funcionalidade do desafio)
- **Três tipos de API**:
  - REST API (HTTP)
  - gRPC
  - GraphQL

## Pré-requisitos

- Docker e Docker Compose

## Instalação e Execução

### Execução Completa com Docker Compose

Para executar todo o sistema (banco de dados, migrações e aplicação) automaticamente:

```bash
docker compose up
```

Este comando irá:
1. Subir o banco de dados MySQL 5.7
2. Subir o RabbitMQ
3. Executar as migrações automaticamente usando golang-migrate
4. Inicializar a aplicação

## Verificação das Migrações

Após executar `docker compose up`, você pode verificar se as migrações foram aplicadas corretamente:

### 1. Verificar Tabelas Criadas

```bash
docker exec -it mysql mysql -uroot -proot -e "USE orders; SHOW TABLES;"
```

**Resultado esperado:**
```
+-------------------+
| Tables_in_orders  |
+-------------------+
| orders            |
| schema_migrations |
+-------------------+
```

### 2. Verificar Estrutura da Tabela Orders

```bash
docker exec -it mysql mysql -uroot -proot -e "USE orders; DESCRIBE orders;"
```

**Resultado esperado:**
```
+-------------+---------------+------+-----+-------------------+-------+
| Field       | Type          | Null | Key | Default           | Extra |
+-------------+---------------+------+-----+-------------------+-------+
| id          | varchar(255)  | NO   | PRI | NULL              |       |
| price       | decimal(10,2) | NO   |     | NULL              |       |
| tax         | decimal(10,2) | NO   |     | NULL              |       |
| final_price | decimal(10,2) | NO   |     | NULL              |       |
| created_at  | timestamp     | NO   |     | CURRENT_TIMESTAMP |       |
+-------------+---------------+------+-----+-------------------+-------+
```

### 3. Verificar Migrações Executadas

```bash
docker exec -it mysql mysql -uroot -proot -e "USE orders; SELECT * FROM schema_migrations;"
```

**Resultado esperado:**
```
+---------+-------+
| version | dirty |
+---------+-------+
|       1 |     0 |
+---------+-------+
```

### 4. Verificar Logs do Migrate

```bash
docker compose logs migrate
```

### 5. Verificar Status dos Containers

```bash
docker compose ps
```

**Resultado esperado:**
- `mysql`: Up e Healthy
- `rabbitmq`: Up e Healthy
- `migrate`: Exited (0) - Completado com sucesso
- `cleanarch-app`: Up

### Troubleshooting de Migrações

Se as migrações não funcionarem:

1. **Verificar conectividade com MySQL:**
   ```bash
   docker exec -it mysql mysql -uroot -proot -e "SELECT 1;"
   ```

2. **Re-executar apenas as migrações:**
   ```bash
   docker compose up migrate
   ```

3. **Verificar logs detalhados:**
   ```bash
   docker compose logs mysql
   docker compose logs migrate
   ```

4. **Reset completo (cuidado - apaga dados):**
   ```bash
   docker compose down -v
   docker compose up
   ```

### Execução em Desenvolvimento Local

#### 1. Subir apenas os serviços de infraestrutura

```bash
docker compose up mysql rabbitmq migrate -d
```

#### 2. Instalar dependências

```bash
go mod tidy
```

#### 3. Gerar código Wire (injeção de dependência)

```bash
cd cmd/ordersystem
go install github.com/google/wire/cmd/wire@latest
wire
```

#### 4. Executar a aplicação localmente

```bash
go run cmd/ordersystem/*.go
```

## Portas dos Serviços

- **REST API**: http://localhost:8000
- **gRPC Server**: localhost:50051
- **GraphQL**: http://localhost:8080
- **GraphQL Playground**: http://localhost:8080 (interface gráfica)
- **MySQL**: localhost:3306
- **RabbitMQ Management**: http://localhost:15672 (guest/guest)

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

#### Listar Orders (Funcionalidade Principal)
```http
GET http://localhost:8000/order
```

### GraphQL

Acesse o GraphQL Playground em: http://localhost:8080

#### Criar Order
```graphql
mutation {
  createOrder(input: {
    id: "order-002"
    price: 150.0
    tax: 15.0
  }) {
    id
    price
    tax
    finalPrice
  }
}
```

#### Listar Orders (Funcionalidade Principal)
```graphql
query {
  orders {
    id
    price
    tax
    finalPrice
  }
}
```

### gRPC

Para testar os serviços gRPC, você pode usar ferramentas como:
- grpcurl
- BloomRPC
- Postman (com suporte gRPC)

#### Exemplo com grpcurl:

```bash
# Criar order
grpcurl -plaintext -d '{"id": "order-003", "price": 200.0, "tax": 20.0}' localhost:50051 pb.OrderService/CreateOrder

# Listar orders (Funcionalidade Principal)
grpcurl -plaintext -d '{}' localhost:50051 pb.OrderService/ListOrders
```

## Arquivo de Testes (api.http)

O arquivo `api.http` contém exemplos de requests para testar todos os endpoints da aplicação. Use com extensões como REST Client no VS Code.

## Migrações do Banco de Dados

As migrações são executadas automaticamente usando golang-migrate:
- **Up**: `sql/migrations/000001_create_orders_table.up.sql`
- **Down**: `sql/migrations/000001_create_orders_table.down.sql`

## Tecnologias Utilizadas

- **Go 1.21**
- **MySQL 5.7**
- **golang-migrate** (migrações)
- **gRPC** com Protocol Buffers
- **GraphQL** com gqlgen
- **REST API** com Gorilla Mux
- **RabbitMQ** (mensageria)
- **Docker & Docker Compose**
- **Wire** (injeção de dependência)

## Estrutura do Projeto (Clean Architecture)

```
cmd/
  ordersystem/          # Ponto de entrada da aplicação
configs/                # Configurações
internal/
  entity/              # Entidades de domínio
  usecase/             # Casos de uso (regras de negócio)
    - create_order.go
    - list_orders.go   # Funcionalidade principal do desafio
  infra/               # Implementações de infraestrutura
    database/          # Repositórios
    web/               # REST API handlers
    grpc/              # Serviços gRPC
    graph/             # Resolvers GraphQL
  event/               # Sistema de eventos
pkg/
  events/              # Event dispatcher
sql/
  migrations/          # Migrações do banco
```

## Testando a Aplicação

1. Execute `docker compose up`
2. Aguarde todos os serviços subirem
3. Use o arquivo `api.http` para testar os endpoints
4. Acesse http://localhost:8080 para o GraphQL Playground
5. Use grpcurl para testar gRPC

## Funcionalidade Principal - Listagem de Orders

O desafio solicitava especificamente a implementação do **usecase de listagem das orders** com:

✅ **Endpoint REST**: `GET /order`  
✅ **Service gRPC**: `ListOrders`  
✅ **Query GraphQL**: `orders`  

Todos implementados e funcionando com MySQL 5.7 e migrações automáticas.
