# DIGIBOOK API

## Description
Digibook API is a simple CRUD REST API for managing book data. This REST API was built with Go programming language and Postgresql as the database.

## Installation
Follow this step to run digibook API on your local machine.
1. Create postgresql database (you can choose the database name)
2. Copy .env.example content to .env file with this command
```sh
cp .env.example .env
```
3. Populate .env file with your config
4. Install all dependencies
```sh
go get .
```
5. Run digibook
```sh
go run .
```

## API Endpoints
- Get all books: `GET` `/api/v1/books`
- Get book by id: `GET` `/api/v1/books/:id`
- Create new book: `POST` `/api/v1/books`
- Update book: `PUT` `/api/v1/books/:id`
- Delete book: `DELETE` `/api/v1/books/:id`