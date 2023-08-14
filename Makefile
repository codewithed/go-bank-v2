postgres:
	docker run --name postgres15 -p 5432:5432  -e POSTGRES_PASSWORD=dbsecret -e POSTGRES_USER=root -d postgres

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go-bank-v2

dropdb:
	docker exec -it postgres15 dropdb --username=root go-bank-v2

migrateup:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/go-bank-v2?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:dbsecret@localhost:5432/go-bank-v2?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go