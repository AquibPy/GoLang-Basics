# Go Module Example

This is a simple example of a Go project using modules. It demonstrates a basic web server using the Gorilla Mux router.

## Getting Started

These instructions will help you set up and run the project on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- A command-line interface (CLI)

### Installation


1. **Initialize the Go module**
   ```sh
   go mod init github.com/yourusername/yourproject
   ```

2. **Download the dependencies**
   ```sh
   go get github.com/gorilla/mux
   ```

3. **Tidy up the dependencies**
   ```sh
   go mod tidy
   ```

### Running the Application

To start the web server, run:
```sh
go run main.go
```

The server will start on port 4000. You can access it by navigating to `http://localhost:4000` in your web browser.

### Code Explanation

Here's a brief explanation of the code:

- **main.go**

    ```go
    package main

    import (
        "fmt"
        "log"
        "net/http"

        "github.com/gorilla/mux"
    )

    func main() {
        fmt.Println("Mod in Go lang")
        greeter()
        r := mux.NewRouter()
        r.HandleFunc("/", serveHome).Methods("GET")
        log.Fatal(http.ListenAndServe(":4000", r))
    }

    func greeter() {
        fmt.Println("Hey there mod users")
    }

    func serveHome(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<h1>Welcome Go Lang</h1>"))
    }
    ```

    - `main()`: The entry point of the application. It prints a welcome message, sets up the router, and starts the server.
    - `greeter()`: A simple function that prints a greeting message.
    - `serveHome()`: Handles HTTP GET requests to the root URL (`/`) and responds with a welcome message.

### Using go mod Commands

- **Initialize a new module**
  ```sh
  go mod init github.com/yourusername/yourproject
  ```

- **Add a dependency**
  ```sh
  go get github.com/gorilla/mux
  ```

- **Tidy up dependencies**
  ```sh
  go mod tidy
  ```

- **Download dependencies**
  ```sh
  go mod download
  ```

- **Verify dependencies**
  ```sh
  go mod verify
  ```


### Acknowledgments

- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for Golang.

```