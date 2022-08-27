# go-finance
Repository for code a finance with Golang

# postgres
- Instalação postgres 14 alpine
   $ docker pull postgres:14-alpine
- Criação do banco
   $ docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
- Execução do postgres 
   $ docker exec -it postgres psql -U postgres

# migrate
- Instalação migrate
    $ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
    $ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
    $ apt-get update
    $ apt-get install -y migrate

- Para executar a criação do migrate. Ex:
    $ migrate create -ext sql -dir db/migrations -seq initial_tables
    - cria dois arquivos (down, up) .sql

- Comandos 
   docker exec -it postgres /bin/sh
   pwd
   ls -l
   createdb --username=postgres --owner=postgres go_finance

# sqlc - Docker
   - docker pull kjconroy/sqlc

Run sqlc using docker run:
   - docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate