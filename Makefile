DB_URL=postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable
postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test