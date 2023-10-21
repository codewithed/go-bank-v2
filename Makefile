postgres:
	docker run --name postgres15 --network bank-network -p 5432:5432  -e POSTGRES_PASSWORD=dbsecret -e POSTGRES_USER=root -d postgres

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-bank-v2

dropdb:
	docker exec -it postgres15 dropdb --username=root go-bank-v2

migrateup:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/go-bank-v2?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/go-bank-v2?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/go-bank-v2?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/go-bank-v2?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination=db/mock/store.go -package=mockdb github.com/codewithed/go-bank-v2/db/sqlc Store 

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1