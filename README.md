# Financial Assistance Scheme Management System

A backend application designed to enable administrators to manage financial assistance schemes and their applications.
This project focuses on creating a backend solution to support individuals and families in need of financial assistance.

The project is built using a hexagonal architecture, ensuring loose coupling for enhanced testability.

It uses [Gin](https://gin-gonic.com/) as the web framework, [PostgreSQL](https://www.postgresql.org/) as the
database, [pgx](https://github.com/jackc/pgx/) as the database driver, along
with [sqlc](https://github.com/sqlc-dev/sqlc) and [Squirrel](https://github.com/Masterminds/squirrel/) as query
builders.

## Table of Contents

- [Usage](#usage)
- [Features](#features)
- [ER Diagram](#er-diagram)
- [License](#license)

## Requirements

1. [Golang-migrate CLI](https://github.com/golang-migrate/migrate)
   - Used for database migrations.

2. [Postgres Database](https://www.postgresql.org/)
   - Used for persistent storage.

## Getting Started

1. Create a Postgres Server using Docker
   ```bash
   docker run --name fas-mgmt-system -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres
   ```
   
2. Run database migrations
   ```bash
   migrate -source file://internal/adapter/storage/postgres/migrations -database postgres://postgres:password@localhost/fas_mgmt_system?sslmode=disable up
   ```
   Edit the database connection details if needed.


3. Create a copy of the .env.example file and rename it to .env

   ```bash
   cp .env.example .env
   ```
   
   Update the values if neccessary.


4. Run the API Server.
   ```bash
   go run .\cmd\api\main.go
   ```

## ER Diagram
![ERD.png](ERD.png)

## API Documentation

The API documentation is located in the `docs/` directory. To access it, open your browser and navigate to
`http://localhost:8080/docs/index.html`. This documentation is generated using [swaggo](https://github.com/swaggo/swag/)
in combination with the [gin-swagger](https://github.com/swaggo/gin-swagger/) middleware.

## License

This project is licensed under the [MIT License](LICENSE).