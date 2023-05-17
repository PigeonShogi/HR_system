createdb:
	docker exec -it postgres_hr createdb --username=root --owner=root hr_system

dropdb: 
	docker exec -it postgres_hr dropdb --username=root hr_system

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hr_system?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hr_system?sslmode=disable" -verbose down

postgres: 
	docker run --name postgres_hr -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb  migrateup migratedown postgres sqlc
