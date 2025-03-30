# Financial Assistance Scheme Management System

Coding assessment for GovTech internship.

A backend application designed to enable administrators to manage financial assistance schemes and their applications.
This project focuses on creating a backend solution to support individuals and families in need of financial assistance.

The project is built using a hexagonal architecture, ensuring loose coupling for enhanced testability.

It uses [Gin](https://gin-gonic.com/) as the web framework, [PostgreSQL](https://www.postgresql.org/) as the
database, [pgx](https://github.com/jackc/pgx/) as the database driver, along
with [sqlc](https://github.com/sqlc-dev/sqlc) and [Squirrel](https://github.com/Masterminds/squirrel/) as query
builders.

## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [File Structure](#file-structure)
- [ER Diagram](#er-diagram)
- [License](#license)

## Requirements

1. [Golang-migrate CLI](https://github.com/golang-migrate/migrate)
   - Used for database migrations.

2. [Docker](https://www.docker.com/)
   - Used for creating a Postgres Server for persistent storage.

## Getting Started

1. Clone this repository
   ```bash
   git clone https://github.com/cxnub/fas-mgmt-system.git
   ```
   
2. Install Dependencies
    ```bash
    go mod download
   ```

1. Create a Postgres Server using Docker
   ```bash
   docker run --name fas-mgmt-system -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres
   ```

2. Create the database

   Use [pgAdmin](https://www.pgadmin.org/) or any other preferred method to connect to the Postgres server and create a database called
   `fas_mgmt_system`.

   
3. Run database migrations
   ```bash
   migrate -source file://internal/adapter/storage/postgres/migrations -database postgres://postgres:password@localhost/fas_mgmt_system?sslmode=disable up
   ```
   Edit the database connection details if needed.


4. Create a copy of the .env.example file and rename it to .env

   ```bash
   cp .env.example .env
   ```
   
   Update the values if neccessary.


5. Run the API Server.
   ```bash
   go run .\cmd\api\main.go
   ```

## API Documentation

The API documentation is located in the `docs/` directory. To access it, open your browser and navigate to
`http://localhost:8080/docs/index.html`. This documentation is generated using [swaggo](https://github.com/swaggo/swag/)
in combination with the [gin-swagger](https://github.com/swaggo/gin-swagger/) middleware.

| Method | Path                                 | Purpose                                                                                                       |
|--------|--------------------------------------|---------------------------------------------------------------------------------------------------------------|
| GET    | /api/applicants                      | Get all applicants.                                                                                           |
| POST   | /api/applicants                      | Create a new applicant.                                                                                       |
| GET    | /api/schemes                         | Get all schemes.                                                                                              |
| GET    | /api/schemes/eligible?applicant={id} | Get all schemes that an applicant (represented by applicant query string parameter) is eligible to apply for. |
| GET    | /api/applications                    | Get all applications.                                                                                         |
| POST   | /api/applications                    | Create a new application.                                                                                     |
| PUT    | /api/applicants/{id}                 | Update an applicant’s details.                                                                                |
| DELETE | /api/applicants/{id}                 | Delete an applicant.                                                                                          |
| POST   | /api/schemes                         | Create a new scheme.                                                                                          |
| PUT    | /api/schemes/{id}                    | Update scheme details.                                                                                        |
| DELETE | /api/schemes/{id}                    | Delete a scheme.                                                                                              |
| PUT    | /api/applications/{id}               | Update application details.                                                                                   |
| DELETE | /api/applications/{id}               | Delete an application.                                                                                        |

Additional routes are displayed in `http://localhost:8080/docs/index.html`. 

   
## File Structure
   ```
   fas-mgmt-system
   ├───cmd
   │   └───api
   ├───docs
   └───internal
       ├───adapter
       │   ├───config
       │   ├───handler
       │   │   └───http
       │   └───storage
       │       └───postgres
       │           ├───migrations
       │           ├───queries
       │           ├───repository
       │           └───sqlc
       └───core
           ├───domain
           ├───port
           ├───service
           └───util

   ```

## ER Diagram
![ERD.png](ERD.png)

## License

This project is licensed under the [MIT License](LICENSE).
