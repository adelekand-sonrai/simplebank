postgres:
	docker run --name postgres-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
		docker exec -it postgres-db createdb --username=root --owner=root simple_bank

dropdb:
		docker exec -it dropdb simple_bank

migrate:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose $(action)

sqlc:
		sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrate sqlc test