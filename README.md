# Expense Tracker

A backend REST API for managing personal expenses, categories, and budgets.

Built with Go and PostgreSQL as a practical backend project.

## Features

* User registration and login
* JWT authentication
* Password hashing with bcrypt
* Category CRUD
* Expense CRUD
* Expense filtering
* Expense statistics
* Budget management
* User data isolation

## Tech Stack

* Go
* Gin
* PostgreSQL
* pgx / pgxpool
* golang-migrate
* JWT
* bcrypt
* Docker Compose

## Architecture

Handler -> Service -> Repository -> PostgreSQL

The project uses a layered architecture to separate HTTP handling, business logic, and database operations.

## Project Structure
cmd/
|-- api/

internal/
|-- apperrors/
|-- config/
|-- database/
|-- models/
|-- dto/
|-- repository/
|-- service/
|-- handler/
|-- token/
|-- middleware/

migrations/

## Run Locally

### Requirements

* Go
* Docker
* Docker Compose
* golang-migrate

### Start PostgreSQL

docker compose up -d

### Run migrations

migrate -path migrations -database "<database-url>" up

### Start the API

go run ./cmd/api/main.go

The API runs on:

http://localhost:8080
