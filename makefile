pullpostgres:
	docker pull postgres:12-alpine

installmigration:
	brew install golang-migrate

migrationcheck:
	migrate -version

postgres:
	docker run --name postgres12_metamaskonline -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12_metamaskonline createdb --username=root --owner=root metamaskonline

dropdb:
	docker exec -it postgres12_metamaskonline dropdb --username=root metamaskonline

dbup:
	migrate -path model/migration -database "postgresql://root:secret@localhost:5432/metamaskonline?sslmode=disable" -verbose up

dbdown:
	migrate -path model/migration -database "postgresql://root:secret@localhost:5432/metamaskonline?sslmode=disable" -verbose down

installsqlc:
	brew install sqlc

sqlc:
	sqlc generate

test:
	go test -v -cover ./...