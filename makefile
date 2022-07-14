pullpostgres:
	docker pull postgres:12-alpine

postgres:
	docker run --name postgres12 --network metamaskonline -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root metamaskonline

dropdb:
	docker exec -it postgres12 dropdb --username=root metamaskonline

