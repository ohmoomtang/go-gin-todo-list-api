# Todo List API

A simple RESTful API for managing a todo list, built with [Gin](https://github.com/gin-gonic/gin) in Go.

## Features

- Create, read, update, and delete todo items
- Simple JSON-based API
- JWT Authorization
- Data storage in MongoDB

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher
- [MongoDB](https://www.mongodb.com/try/download/community) (local, Docker, or cloud)

## Installation

```bash
git clone https://github.com/yourusername/go-gin-todo-list-api.git
cd go-gin-todo-list-api
go mod tidy
```

## Configuration

Set the following environment variables before running the API:

- `MONGODB_URI`: MongoDB connection URI
- `JWT_KEY`: Secret key for signing/verifying JWT tokens

## Running the API

```bash
go run main.go
```

The API will be available at [http://localhost:8080](http://localhost:8080).

## API Endpoints

| Method | Endpoint        | Description             |
|--------|----------------|------------------------|
| POST   | /register      | Register new user      |
| GET    | /login         | Login (get JWT Token)  |
| GET    | /todos         | List all todos         |
| GET    | /todos/:id     | Get a todo by ID       |
| POST   | /todos         | Create a new todo      |
| PUT    | /todos/:id     | Update a todo by ID    |
| DELETE | /todos/:id     | Delete a todo by ID    |

## Example Todo Item

```json
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Buy milk, eggs, and bread"
}
```

## License

MIT