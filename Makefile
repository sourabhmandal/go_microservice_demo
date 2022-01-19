postgres:
	docker stop postgres12 && docker rm postgres12 && docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb test