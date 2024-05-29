postgres:
	@echo "initializing postgres..."
	docker run --name postgres_go_bank --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine3.19

createdb:
	@echo "creating db..."
	docker exec -it postgres_go_bank createdb --username=root --owner=root go_bank

dropdb:
	@echo "dropping db..."
	docker exec -it postgres_go_bank dropdb go_bank


migrate_up:
	@echo "running up migrations..."
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose up


migrate_down:
	@echo "running down migrations..."
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose down

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
start: postgres createdb migrate_up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/henryeffiong/gobank/db/sqlc Store

proto:
	rm -f pb/*.proto
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto


.PHONY: postgres createdb dropdb migrate_up migrate_down sqlc test server mock proto

# migrate create -ext sql -dir db/migration -seq NAME