# API Bank Transactions

## About the project

- An api for transferring values to wallets.
A user creates his account, then creates a wallet, so he can receive and send transfers.
App use Golang, RabbitMQ, Postgres

## How to run the api?

**Step 1: Clone the project, run the following commands:**

**Step 2: Create file config.json**
{
  "database": {
    "url": "://postgres:postgres@172.22.0.2:5432/db_bank"
  },
  "server": {
    "port": "3001"
  }
}

**Step 3: Run docker(at the root of the project)**
- docker-composer up -d

**Step 4: Create database:**
- Create database: db_bank <br>

**Step 5: go mod tidy(at the root of the project)**
- go mod tidy

# Link Doc(Swagger)
http://localhost:3001/swagger/index.html

# Link RabbitMQ
http://localhost:15672/
Login: guest
Password: guest

Create one queue with name (transactions)
Add (amq.direct) in Bindings/From exchange

# Link Grafana
http://localhost:3000/
Login: admin
Password: admin

Click in Data Sources and Prometheus for configuration.
change url for; http://prometheus:9090 and save.
Copy ID Dashboard in; https://grafana.com/grafana/dashboards/10991-rabbitmq-overview/
Import Dashboard in Grafana, using ID.

# Link Flowchart
https://whimsical.com/registra-em-transacoes-e-nao-credita-na-conta-do-usuario-7dgWckZDKij8GYTNyk4XVK


