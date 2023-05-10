createdb:
	docker exec -it postgres_hr createdb --username=root --owner=root hr_system

dropdb: 
	docker exec -it postgres_hr dropdb --username=root hr_system

postgres: 
	docker run --name postgres_hr -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

.PHONY: postgres
