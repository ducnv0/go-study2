# 4-Hour Python-to-Go Crash Course

This is an intensive, hands-on 4-hour plan to get a backend Python developer up and running with Go. Because you already understand APIs and logic concepts, we are skipping the basics of programming and focusing purely on syntax, structure, and the "Go Way".

*Goal:* By the end of 4 hours, you will understand Go syntax, struct-based modeling, writing a web API, and running concurrent background tasks.

---

## Hour 1: The Basics (Syntax & Mechanics)
**Objective:** Get comfortable reading Go code and understand how it differs from Python.

> [!NOTE]
> Open your terminal and create a new folder for this hour: `mkdir crash-course && cd crash-course`

* **00:00 - 00:15: Setup & "Hello World"**
  * Initialize a project module: `go mod init crash-course`
    * *Note on `go.mod`: This completely replaces Python's `venv`, `pip`, and `requirements.txt`. It defines the project root and tracks exact dependency versions. Packages are stored in a single global cache (`~/go/pkg/mod`), completely eliminating the need for virtual environments!*
  * Write `main.go` using `package main` and `func main() { fmt.Println("Hello") }`.
  * **Execution Differences**:
    * `go run main.go`: Compiles your code in memory and executes it immediately. It's meant for local development and feels exactly like running `python main.py`.
    * `go build`: Compiles your project into a permanent, highly-optimized standalone binary file on your disk. You can take this resulting file (e.g., `./crash-course`) and run it on a production server *without even needing Go installed!*
* **00:15 - 00:30: Variables & Types**
  * Try Type Inference inside functions: `name := "Alice"` (Go figures out it's a string, acts like `name = "Alice"`).
  * Declare explicitly: `var age int = 30`.
* **00:30 - 00:45: Collections (Lists & Dicts)**
  * **Slices** (your Python Arrays/Lists): `nums := []int{1, 2, 3}`. Practice adding to a slice: `nums = append(nums, 4)`.
  * **Maps** (your Python Dicts): `user := map[string]string{"name": "Alice"}`. Retrieve a value: `name := user["name"]`.
* **00:45 - 01:00: Control Flow & Error Handling**
  * Go has NO `while` loop. The `for` loop does everything.
  * Practice looping over your slice: `for index, value := range nums { ... }`.
  * The Python `try/except` replacement: Call a function that returns an error, and immediately check `if err != nil { fmt.Println(err) }`.

## Hour 2: Data Modeling (The "Object-Oriented" Part)
**Objective:** Move away from Python classes, `self`, and `__init__`. Comprehend Structs and Pointers.

* **01:00 - 01:20: Pointers (Demystified)**
  * Understand `&` (get memory address) and `*` (read/write to that address). 
  * Passing a pointer in Go is exactly like passing a dictionary into a function in Python—it modifies the original object payload rather than a duplicate.
* **01:20 - 01:40: Structs & Methods (Python `class` equivalent)**
  * Define a model: `type User struct { Name string; Age int }`.
  * Attach a method: `func (u *User) SayHello() { fmt.Println(u.Name) }`. Notice that `(u *User)` is the *receiver* and it completely replaces Python's `self`.
* **01:40 - 02:00: Struct Tags (Pydantic Equivalent)**
  * Learn how Go maps structs to JSON. Define `User` with tags: ``type User struct { Name string `json:"user_name"` }``.
  * Use the standard `encoding/json` package. Call `json.Marshal(user)` and print the output as a string.

## Hour 3: The Web API (FastAPI Equivalent)
**Objective:** Build an HTTP server, expose routing endpoints, and validate input payloads.

* **02:00 - 02:20: The Standard Library Server**
  * Use `net/http` to build a simple server: `http.HandleFunc("/hello", myFunc)` and `http.ListenAndServe(":8080", nil)`. Hit it with `curl` to see it working.
* **02:20 - 02:40: Introduction to Gin (High-Speed Router)**
  * Install the popular framework: `go get -u github.com/gin-gonic/gin`.
  * Start up a server with `r := gin.Default()` and write a `r.GET("/ping")` route returning JSON: `c.JSON(200, gin.H{"msg": "pong"})`.
* **02:40 - 03:00: A POST Request Endpoint**
  * Create a `r.POST("/users")` endpoint. Create a payload struct `type CreateUserRequest struct { ... }`.
  * Use Gin's `c.ShouldBindJSON(&payload)` to validate the incoming request body payload (this acts exactly like an incoming Pydantic BaseModel in FastAPI).

## Hour 4: Concurrency Unleashed (Goodbye `asyncio`)
**Objective:** Use Go's biggest superpower—Goroutines. Achieve parallelism effortlessly without managing event loops.

* **03:00 - 03:20: Goroutines**
  * Write a slow function using `time.Sleep(2 * time.Second)`. 
  * Run it natively in the background by just prepending the `go` keyword: `go mySlowFunction()`. (Notice the program will exit immediately if `main()` finishes before your function does).
* **03:20 - 03:40: Channels / WaitGroups (Safe data passing)**
  * To make `main()` wait for the background task to finish, import `"sync"` and use `wg := sync.WaitGroup{}` (call `wg.Add(1)`, `wg.Done()`, and `wg.Wait()`).
  * Create a channel: `ch := make(chan string)`. Send data into it from the background function: `ch <- "Finished Work!"`. Receive it in `main()`: `msg := <-ch`.
* **03:40 - 04:00: Final Challenge (Concurrent Fetcher)**
  * Write a script that takes a slice of 5 different URLs. Use a `for` loop to spin up a goroutine for each one. Have each goroutine make a quick `http.Get()` call. Use a `sync.WaitGroup` to wait for all downloads to complete, and print how long the whole script took. You'll be amazed by the speed!

---

## Quick Prep Checklist
Do this *before* you start the clock:
1. **Tooling**: Ensure you have installed Go. Ensure you are using **VS Code with the Go extension** (or GoLand). The IDE automatically format your code and injects imports on save, which is crucial for moving fast.
2. **Terminal Ready**: Keep an empty terminal window open specifically for `go run` and `go get`.
