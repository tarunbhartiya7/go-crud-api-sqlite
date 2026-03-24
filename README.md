# Events CRUD API

Event management REST API built with Go, Gin, and SQLite, featuring JWT-based authentication, user signup/login with bcrypt password hashing, full CRUD for events, protected event registration/cancellation routes with ownership and validation checks, and integrated Swagger/OpenAPI documentation for interactive API exploration.

## Features

- **Events CRUD** — Create, read, update, and delete events
- **User signup & login** — Register and login with JWT authentication
- **Password hashing** — bcrypt for secure password storage
- **Relational data** — Events linked to users via foreign key
- **Event registration** — Register/cancel event participation for authenticated users
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

## Swagger Docs

Generate OpenAPI docs:

```bash
go run github.com/swaggo/swag/cmd/swag@latest init -g main.go
```

Run the API and open Swagger UI:

- `http://localhost:8080/swagger/index.html`

For protected endpoints, set the `Authorization` header in Swagger UI to your JWT token (without `Bearer` prefix in this project).

## API Endpoints

### Events

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /events | List all events |
| GET | /events/:id | Get event by ID |
| POST | /events | Create event |
| PUT | /events/:id | Update event |
| DELETE | /events/:id | Delete event |
| POST | /events/:id/register | Register current user for an event |
| DELETE | /events/:id/register | Cancel registration for an event |

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

Requires `Authorization` header with JWT token from login.

```bash
curl -X PUT http://localhost:8080/events/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: YOUR_JWT_TOKEN" \
  -d '{"name":"Updated Concert","description":"Updated description","location":"New Venue"}'
```

**Response:** `200 OK` — Updated event  
**Error:** `404 Not Found` — Event does not exist  
**Error:** `403 Forbidden` — You are not the owner of this event

### DELETE /events/:id

Requires `Authorization` header with JWT token from login.

```bash
curl -X DELETE http://localhost:8080/events/1 \
  -H "Authorization: YOUR_JWT_TOKEN"
```

**Response:** `200 OK` — `{"message":"Event deleted successfully"}`  
**Error:** `404 Not Found` — Event does not exist  
**Error:** `403 Forbidden` — You are not the owner of this event

### POST /events/:id/register

Requires `Authorization` header with JWT token from login.

```bash
curl -X POST http://localhost:8080/events/1/register \
  -H "Authorization: YOUR_JWT_TOKEN"
```

**Response:** `201 Created` — `{"message":"Registered for event"}`  
**Error:** `404 Not Found` — Event does not exist

### DELETE /events/:id/register

Requires `Authorization` header with JWT token from login.

```bash
curl -X DELETE http://localhost:8080/events/1/register \
  -H "Authorization: YOUR_JWT_TOKEN"
```

**Response:** `200 OK` — `{"message":"Registration cancelled"}`  
**Error:** `404 Not Found` — Registration not found

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
curl -X POST http://localhost:8080/signup -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"secret123"}' | jq
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"secret123"}' | jq

# Events (replace TOKEN with JWT from login)
curl http://localhost:8080/events
curl http://localhost:8080/events/1
curl -X POST http://localhost:8080/events -H "Content-Type: application/json" -H "Authorization: TOKEN" -d '{"name":"Concert","description":"Live music","location":"Central Park"}' | jq
curl -X PUT http://localhost:8080/events/7 -H "Content-Type: application/json" -H "Authorization: TOKEN" -d '{"name":"Updated","description":"Updated","location":"New Venue"}' | jq
curl -X DELETE http://localhost:8080/events/7 -H "Authorization: TOKEN" | jq
curl -X POST http://localhost:8080/events/1/register -H "Authorization: TOKEN" | jq
curl -X DELETE http://localhost:8080/events/1/register -H "Authorization: TOKEN" | jq
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
│   ├── register.go   # Event registration handlers
│   └── users.go      # User handlers
├── docs/             # Generated Swagger docs (swagger.json/yaml)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── middlewares/
│   └── auth.go       # JWT auth middleware
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
- [swaggo](https://github.com/swaggo/swag) — OpenAPI generation and Swagger UI
