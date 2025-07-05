SHELL := /bin/bash
DB_URL = postgresql://root:secret@localhost:5432/go_simple_bank?sslmode=disable

postgres:
	sudo docker run --name postgres-sb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.9-bullseye
createdb:
	sudo docker exec -it postgres-sb createdb --username=root --owner=root go_simple_bank
dropdb:
	sudo docker exec -it postgres-sb dropdb go_simple_bank
migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up
migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down
migrateforce-1:
	migrate -path db/migration -database "${DB_URL}" -verbose force 1
sqlc-g:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	source ~/.zshrc && mockgen -package mockdb -destination db/mocke/store.go github.com/john9101/go-simplebank/db/sqlc Store
proto:
	rm -f pb/*.go
	PATH=$(PATH):$(HOME)/go/bin \
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto
evans:
	evans --port 9090 --host localhost --package pb --service GoSimpleBank -r repl
.PHONY: postgres createdb dropdb migrateup migratedown migrateforce-1 sqlc-g test server mock proto evans
