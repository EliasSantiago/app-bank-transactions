# app-bank-transactions
App for bank transactions with Golang, RabbitMQ, Postgres...

# Link Doc(Swagger)
http://localhost:3001/swagger/index.html

# Link Flowchart
https://whimsical.com/registra-em-transacoes-e-nao-credita-na-conta-do-usuario-7dgWckZDKij8GYTNyk4XVK

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


