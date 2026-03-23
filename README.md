# Events CRUD API

Event management API with user signup, bcrypt hashing, and SQLite storage — built with Go and Gin.

## Features

- **Events CRUD** — Create, read, update, and delete events
- **User signup & login** — Register and login with JWT authentication
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
| POST | /login | Login and get JWT token |

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

Requires `Authorization` header with JWT token from login.

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -H "Authorization: YOUR_JWT_TOKEN" \
  -d '{"name":"Concert","description":"Live music performance","location":"Central Park"}'
```

`dateTime` and `userId` are set by the server from the token.

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

### POST /login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'
```

**Response:** `200 OK` — `{"message":"Login successful","token":"eyJhbGc..."}`  
**Error:** `401 Unauthorized` — Invalid credentials

## cURL Quick Reference

```bash
# Auth
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"secret123"}'
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"secret123"}'

# Events (replace TOKEN with JWT from login)
curl http://localhost:8080/events
curl http://localhost:8080/events/1
curl -X POST http://localhost:8080/events -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXI1QGV4YW1wbGUuY29tIiwiZXhwIjoxNzc0Mjc1NzIyLCJ1c2VySWQiOjZ9.K11IOOTwpb7G00GrRNWJTEV0t0RB8Phi4J0vWePXjXU" -d '{"name":"Concert","description":"Live music","location":"Central Park"}'
curl -X PUT http://localhost:8080/events/1 -H "Content-Type: application/json" -d '{"name":"Updated","description":"Updated","location":"New Venue"}'
curl -X DELETE http://localhost:8080/events/1
```

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
│   ├── hash.go       # bcrypt password hashing
│   └── jwt.go        # JWT generation and verification
└── events.db         # SQLite database (created on first run)
```

## Tech Stack

- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [go-sqlite3](https://github.com/mattn/go-sqlite3) — SQLite driver for Go
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — Password hashing
- [jwt-go](https://github.com/golang-jwt/jwt) — JWT authentication
