# Events CRUD API

Event management API with user signup, bcrypt hashing, and SQLite storage — built with Go and Gin.

## Features

- **Events CRUD** — Create, read, update, and delete events
- **User signup** — Register users with email and password
- **Password hashing** — bcrypt for secure password storage
- **Relational data** — Events linked to users via foreign key
- **Error handling** — 404 for missing resources, validation for invalid input
- **Security** — Passwords never returned in API responses

## Prerequisites

- Go 1.25+

## Getting Started

```bash
# Install dependencies
go mod download

# Run the server
go run main.go
```

The server starts at `http://localhost:8080`.

## API Endpoints

### Events

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /events | List all events |
| GET | /events/:id | Get event by ID |
| POST | /events | Create event |
| PUT | /events/:id | Update event |
| DELETE | /events/:id | Delete event |

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /signup | Register new user |

### GET /events

```bash
curl http://localhost:8080/events
```

**Response:** `200 OK` — Array of events

### GET /events/:id

```bash
curl http://localhost:8080/events/1
```

**Response:** `200 OK` — Event object  
**Error:** `404 Not Found` — Event does not exist

### POST /events

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{"name":"Concert","description":"Live music performance","location":"Central Park"}'
```

`dateTime` and `userId` are set by the server.

**Response:** `201 Created` — Created event  
**Error:** `400 Bad Request` — Invalid or missing fields

### PUT /events/:id

```bash
curl -X PUT http://localhost:8080/events/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Concert","description":"Updated description","location":"New Venue"}'
```

**Response:** `200 OK` — Updated event  
**Error:** `404 Not Found` — Event does not exist

### DELETE /events/:id

```bash
curl -X DELETE http://localhost:8080/events/1
```

**Response:** `200 OK` — `{"message":"Event deleted successfully"}`  
**Error:** `404 Not Found` — Event does not exist

### POST /signup

```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'
```

**Response:** `200 OK` — `{"message":"User created successfully","user":{"id":1,"email":"user@example.com"}}`  
**Error:** `400 Bad Request` — Invalid or missing fields

## Project Structure

```
.
├── main.go           # Entry point, DB init, route setup
├── db/
│   └── db.go         # SQLite connection and schema
├── models/
│   ├── event.go      # Event model and data access
│   └── user.go       # User model and data access
├── routes/
│   ├── routes.go     # Route registration
│   ├── events.go     # Event handlers
│   └── users.go      # User handlers
├── utils/
│   └── hash.go       # bcrypt password hashing
└── events.db         # SQLite database (created on first run)
```

## Tech Stack

- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [go-sqlite3](https://github.com/mattn/go-sqlite3) — SQLite driver for Go
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — Password hashing
