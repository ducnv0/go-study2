# Fast-Track Go Learning Plan for Python Developers

Welcome to the Go ecosystem! As a Python dev (Django/FastAPI), you already know how to build robust APIs and handle complex data. Moving to Go (Golang) isn't about learning *how* to build backends again; it's about shifting your mindset regarding typing, project structure, and concurrency.

This document serves as your map, translator, and 4-week action plan.

---

## 1. The Paradigm Shift (Python vs. Go)

Go forces a slightly different mental model compared to Python:

> [!NOTE]
> **Compiled & Statically Typed**: Your code compiles to a single binary. Errors are caught *before* you run the code. There's no interpreter, making execution significantly faster and deployments beautifully simple (no `gunicorn`, no `venv`).

> [!IMPORTANT]
> **Composition over Inheritance**: Go does not have `class`, `extends`, or `super()`. You will use `struct`s to hold data and attach methods to them. You use `interface`s to define behavior.

> [!TIP]
> **Explicit Error Handling**: No `try...except`. Functions that can fail return an `error` as their last return value. You will write `if err != nil` a lot, and this is considered a feature, not a bug, making control flow predictable.

> [!TIP]
> **Built-in Concurrency**: Forget the Global Interpreter Lock (GIL), `asyncio`, and `celery` for simple background tasks. Go uses **Goroutines** (lightweight virtual threads managed by the Go runtime) and **Channels** (for goroutines to talk to each other safely).

---

## 2. The Python to Go Rosetta Stone

### Basics & Package Management
*   **Python**: `pip`, `venv`, `requirements.txt`
*   **Go**: `go mod`. (Run `go mod init <module-name>`). It automatically tracks dependencies in `go.mod` and `go.sum`. 

### Variables & Instantiation
```python
# Python
name: str = "Alice"
user = User(name="Alice")
```
```go
// Go
name := "Alice" // Type explicitly inferred
user := User{Name: "Alice"} 
```

### Functions & Multiple Returns
```python
# Python
def divide(a: int, b: int) -> tuple[float, Exception | None]:
    if b == 0:
        return 0, Exception("division by zero")
    return a / b, None
```
```go
// Go
func divide(a int, b int) (int, error) { // Multiple return types
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Object-Oriented Programming (Classes vs Structs)
```python
# Python
class User:
    def __init__(self, name: str):
        self.name = name

    def say_hello(self):
        print(f"Hello, {self.name}")
```
```go
// Go
type User struct {
    Name string // Capitalized means 'public' (exported) outside the package
}

// Method receiver attached to a pointer of User
func (u *User) SayHello() {
    fmt.Printf("Hello, %s\n", u.Name) // string interpolation
}
```

---

## 3. The Backend Stack Equivalents

What do you use when replacing your Python tools in Go?

| Python Concept | Go World Equivalent | Notes |
| :--- | :--- | :--- |
| **FastAPI / Django** | `net/http` (stdlib), **Gin**, **Echo**, **Fiber** | The standard library is incredibly powerful. However, **Gin** is very similar to FastAPI in speed, routing capability, and ease of use. |
| **SQLAlchemy / Django ORM** | **GORM**, **sqlc**, `database/sql` | GORM is the closest to SQLAlchemy. However, a major trend in Go is **sqlc**, which *generates type-safe Go code directly from raw SQL queries*. |
| **Pydantic** | **Struct Tags**, `go-playground/validator` | You define JSON validation right against your struct definitions (e.g., ``Name string `json:"name" binding:"required"` ``). |
| **Celery / RQ** | **Goroutines**, **Asynq**, **Machinery** | For 80% of tasks, a simple `go backgroundTask()` is enough. For persistent, retriable queues with Redis, **Asynq** is standard. |
| **Pytest** | `testing` package | Built natively (`go test`). You name test files `file_name_test.go`. |
| **Uvicorn / Gunicorn** | *Built-in* | The compiled Go binary starts an incredibly highly-concurrent web server natively. Just execute `./myapp`. |

---

## 4. Your Accelerated Action Plan (4 Weeks)

### Week 1: Syntax & Core Mechanics
1. **Learn**: Go through [A Tour of Go](https://go.dev/tour/) fully. This teaches syntax, loops, slices, and maps (Go's dicts).
2. **Understand Pointers**: Don't be intimidated! Think of them as references you pass so you don't copy heavy objects around (`&` gets the memory address, `*` reads the value at the address).
3. **Practice**: Create a CLI tool in Go. Use `net/http` to do a GET request to a public API (like PokeAPI or Weather), and unmarshal the JSON into Go structs using `encoding/json`.

### Week 2: Build Your First API (The FastAPI Way)
1. **Framework Selection**: Start with **Gin** (`github.com/gin-gonic/gin`) or just stick to the standard library with `go-chi/chi`.
2. **Database Connect**: Boot a Postgres database. Learn how to connect using `jackc/pgx` (the standard PostgreSQL driver in Go). 
3. **Task**: Rebuild a simple CRUD API you once built in Python (e.g., an Articles or Users API). Write the struct models, handle JSON bindings, and save data to DB.

### Week 3: Project Structure & Interfaces
1. **Architecture**: Read about standard Go layout. Group code by domain (e.g., putting all User logic and routers in an `internal/user` package). Avoid circular dependencies.
2. **Interfaces**: In Go, you accept interfaces and return structs. Learn how to mock a database service for testing using interfaces.
3. **Task**: Refactor your CRUD API to inject the database connection into your HTTP handlers, rather than defining the DB connection as a global variable. Add table-driven tests (`_test.go`).

### Week 4: Concurrency Unleashed
1. **Goroutines**: Add a background job to your API. E.g., when a User registers, fire off a `go sendWelcomeEmail(user)` function so the HTTP response returns immediately without blocking.
2. **Channels & WaitGroups**: Learn `sync.WaitGroup` to wait for 10 concurrent requests to finish before proceeding.
3. **Deployment**: Write a tiny multi-stage `Dockerfile`. You'll be amazed when your fully functional web API Docker image is <20 MB and uses purely 15MB of RAM at runtime!

---

## 5. Must-Read Resources & Docs
* [Effective Go](https://go.dev/doc/effective_go) - Once you know the syntax, this is mandatory. It explains how to write *idiomatic* Go.
* [Go by Example](https://gobyexample.com/) - A hands-on introduction to Go using annotated example programs.
* [Let's Go](https://lets-go.alexedwards.net/) / [Let's Go Further](https://lets-go-further.alexedwards.net/) - These are widely considered the masterclass books for building production web APIs in Go.
