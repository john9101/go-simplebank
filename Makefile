SHELL := /bin/bash

postgres:
	sudo docker run --name postgres-sb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.9-bullseye
createdb:
	sudo docker exec -it postgres-sb createdb --username=root --owner=root simple_bank
dropdb:
	sudo docker exec -it postgres-sb dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrateforce-1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose force 1
sqlc-g:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	source ~/.zshrc && mockgen -package mockdb -destination db/mocke/store.go github.com/john9101/simplebank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown migrateforce-1 sqlc-g test server mock