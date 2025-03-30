.PHONY: postgres adminer migrate

postgres:
	docker run --name fas-mgmt-system -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres

sqlc_gen:
	docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate

migrate up:
 	migrate -source file://internal/adapter/storage/postgres/migrations -database postgres://postgres:password@localhost/fas_mgmt_system?sslmode=disable up

migrate down:
 	migrate -source file://internal/adapter/storage/postgres/migrations -database postgres://postgres:password@localhost/fas_mgmt_system?sslmode=disable down


start:
	go run cmd/api/main.go