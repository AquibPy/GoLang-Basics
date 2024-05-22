# Netflix Watchlist API

This is a Go API for managing a Netflix watchlist, using MongoDB as the database. The API supports CRUD operations and is documented with Swagger.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Swagger Integration](#swagger-integration)
- [Running the Application](#running-the-application)

## Features

- Create, Read, Update, and Delete movies in a MongoDB database
- Mark movies as watched
- Delete all movies
- Swagger documentation for API endpoints

## Installation

### Prerequisites

- Go (version 1.16 or later)
- MongoDB
- Git

### Steps

1. Install the required Go packages:

```sh
go mod tidy
```

2. Install `swag` for generating Swagger documentation:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

3. Ensure `swag` is available in your PATH:

    - **Unix/Linux/MacOS**: Add the following line to your `~/.bashrc`, `~/.bash_profile`, or `~/.zshrc` file (depending on your shell):

        ```sh
        export PATH=$(go env GOPATH)/bin:$PATH
        ```

    - **Windows**: Add `%GOPATH%\bin` to your system's PATH.

## Usage

### MongoDB Connection

Update the `connectionString` constant in `controller/controller.go` with your MongoDB connection string:

```go
const connectionString = "mongodb+srv://<username>:<password>@cluster0.sdx7i86.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
```

Replace `<username>` and `<password>` with your MongoDB credentials.

### Generating Swagger Documentation

Run the following command to generate Swagger documentation:

```sh
swag init
```

This will create two files: `docs/docs.go` and `docs/swagger.json`.

## API Endpoints

### Get All Movies

- **URL**: `/api/movies`
- **Method**: `GET`
- **Description**: Retrieves all movies from the watchlist.

### Create a Movie

- **URL**: `/api/movie`
- **Method**: `POST`
- **Description**: Adds a new movie to the watchlist.
- **Request Body**:
  ```json
  {
    "movie": "string",
    "watched": "boolean"
  }
  ```

### Mark a Movie as Watched

- **URL**: `/api/movie/{id}`
- **Method**: `PUT`
- **Description**: Updates a movie's watched status to true.
- **Path Parameter**: `id` (string)

### Delete a Movie

- **URL**: `/api/movie/{id}`
- **Method**: `DELETE`
- **Description**: Deletes a movie from the watchlist by ID.
- **Path Parameter**: `id` (string)

### Delete All Movies

- **URL**: `/api/deleteallmovie`
- **Method**: `DELETE`
- **Description**: Deletes all movies from the watchlist.

## Swagger Integration

Swagger is used to document the API. Follow these steps to set it up:

1. **Install Swagger**: Ensure you have installed `swag`:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

2. **Add Swagger Annotations**: Annotate your handler functions in `controller/controller.go` with Swagger comments. Here is an example for the `GetMyAllMovies` function:

```go
// @Summary Get all movies
// @Description Get all movies from the watchlist
// @Tags movies
// @Produce json
// @Success 200 {array} model.Netflix
// @Router /api/movies [get]
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    allMovies := getAllMovies()
    json.NewEncoder(w).Encode(allMovies)
}
```

3. **Generate Swagger Docs**: Run the command `swag init` in your project directory to generate the Swagger documentation files (`docs/docs.go` and `docs/swagger.json`).

4. **Serve Swagger UI**: Modify `main.go` to serve the Swagger UI:

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/AquibPy/mongoapi/router"
    httpSwagger "github.com/swaggo/http-swagger"
    _ "github.com/AquibPy/mongoapi/docs" // Import the generated docs
)

func main() {
    fmt.Println("MongoDB API")
    r := router.Router()
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    fmt.Println("Server is getting started....")
    log.Fatal(http.ListenAndServe(":4000", r))
    fmt.Println("Listening at port 4000 ...")
}
```

5. **Access Swagger UI**: Open your browser and navigate to `http://localhost:4000/swagger/index.html` to view and interact with the Swagger UI.

## Running the Application

Run the following command to start the application:

```sh
go run main.go
```

The server will start on port 4000. You can interact with the API using any API client (e.g., Postman) or through the Swagger UI at `http://localhost:4000/swagger/index.html`.
