# Events CRUD API

A simple REST API for managing events, built with Go, Gin, and SQLite.

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

### GET /events

List all events.

```bash
curl http://localhost:8080/events
```

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "name": "Event 1",
    "description": "Description 1",
    "location": "Location 1",
    "dateTime": "2025-03-21T12:00:00Z",
    "userId": 1
  }
]
```

### GET /events/:id

Get a single event by ID.

```bash
curl http://localhost:8080/events/1
```

**Response:** `200 OK` — Returns the event object  
**Error:** `404 Not Found` — Event does not exist

### POST /events

Create a new event. `dateTime` and `userId` are set by the server.

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{"name":"Concert","description":"Live music performance","location":"Central Park"}'
```

**Request body:**

```json
{
  "name": "Concert",
  "description": "Live music performance",
  "location": "Central Park"
}
```

**Response:** `201 Created` — Returns the created event with `id`, `dateTime`, and `userId`  
**Error:** `400 Bad Request` — Invalid JSON or missing required fields

### PUT /events/:id

Update an existing event.

```bash
curl -X PUT http://localhost:8080/events/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Concert","description":"Updated description","location":"New Venue"}'
```

**Response:** `200 OK` — Returns the updated event  
**Error:** `404 Not Found` — Event does not exist

### DELETE /events/:id

Delete an event.

```bash
curl -X DELETE http://localhost:8080/events/1
```

**Response:** `200 OK` — `{"message":"Event deleted successfully"}`  
**Error:** `404 Not Found` — Event does not exist

## Project Structure

```
.
├── main.go           # Entry point, DB init, route setup
├── db/
│   └── db.go         # SQLite connection and schema
├── models/
│   └── event.go      # Event model and data access
├── routes/
│   ├── routes.go     # Route registration
│   └── events.go     # Event handlers
└── events.db         # SQLite database (created on first run)
```

## Tech Stack

- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [go-sqlite3](https://github.com/mattn/go-sqlite3) — SQLite driver for Go
