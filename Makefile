postgres:
	docker run --name postgres -p 5423:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5423/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5423/simple_bank?sslmode=disable" -verbose up 1
	
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5423/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5423/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/zura-t/simplebank/db/sqlc Store

.PHONY: postgres test sqlc createdb dropdb mock migratedown migrateup migratedown2 migrateup1 server