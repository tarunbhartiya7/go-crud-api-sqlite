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

### POST /events

Create a new event.

**Request body:**

```json
{
  "name": "Concert",
  "description": "Live music performance",
  "location": "Central Park",
  "dateTime": "2025-04-15T18:00:00Z",
  "userId": 1
}
```

**Response:** `201 Created`

## Project Structure

```
.
├── main.go       # Server setup and route handlers
├── db/
│   └── db.go     # SQLite connection and schema
├── models/
│   └── event.go  # Event model and business logic
└── events.db     # SQLite database (created on first run)
```

## Tech Stack

- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [go-sqlite3](https://github.com/mattn/go-sqlite3) — SQLite driver for Go
