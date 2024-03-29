postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine 

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres go_finance

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

sqlc-gen:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb migrateup migrationdrop test server sqlc-gen