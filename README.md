# Go TODO REST API

A simple TODO REST API built in Go to learn the language fundamentals
and web service development patterns.

Built as part of my learning journey transitioning from Java/Spring Boot to Go.

## Tech Stack
- Go 1.24
- net/http (built-in HTTP server)
- encoding/json (built-in JSON handling)
- In-memory storage (PostgreSQL coming soon!)

## Endpoints
| Method | URL | Description |
|--------|-----|-------------|
| GET | /todos | Get all todos |
| POST | /todos/create | Create a new todo |
| PUT | /todos/update/{id} | Update a todo by ID |
| DELETE | /todos/{id} | Delete a todo by ID |

## What I Learned
- Go syntax and language fundamentals coming from Java/Spring Boot
- HTTP server setup with net/http package
- JSON encoding/decoding with encoding/json
- RESTful API patterns in Go
- Struct-based data modelling vs Java classes
- Go error handling vs Java try/catch
- Slice manipulation for in-memory CRUD operations

## How Go Compares to Java (My Perspective)
- No classes — use structs instead
- No try/catch — explicit error return values
- No Spring annotations — routes registered centrally in main()
- Much less boilerplate — no public/private keywords needed
- Built-in HTTP server — no Tomcat/embedded server needed

## Running Locally
```bash
git clone https://github.com/Brendanhwz/go-todo-api.git
cd go-todo-api
go run main.go
# Server runs on localhost:8080
```

## Testing with Postman
Import these requests into Postman:

**Create a todo:**
- Method: POST
- URL: http://localhost:8080/todos/create
- Body: `{"title":"Learn Golang","done":false}`

**Get all todos:**
- Method: GET
- URL: http://localhost:8080/todos

**Update a todo:**
- Method: PUT
- URL: http://localhost:8080/todos/update/1
- Body: `{"title":"Learn Golang - UPDATED","done":true}`

**Delete a todo:**
- Method: DELETE
- URL: http://localhost:8080/todos/1

## Upcoming Enhancements
- [ ] PostgreSQL database integration
- [ ] Middleware (logging + CORS)
- [ ] Better error handling
- [ ] gorilla/mux router

## Author
Brendan | Software Engineer
Transitioning from Java/Spring Boot to Go