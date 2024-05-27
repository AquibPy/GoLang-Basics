# URL Shortener Service

This project is a simple URL shortener service written in Go. It allows users to shorten a long URL into a short one, and later redirect to the original URL using the short version.

## Features

- Shorten a given URL.
- Redirect to the original URL using the shortened version.
- JSON-based API responses.

## Getting Started

### Prerequisites

- Go 1.16 or later


The server will start on port 3000.

## API Endpoints

### Welcome Endpoint

- **URL:** `/`
- **Method:** `GET`
- **Description:** Returns a welcome message.
- **Response:**
  ```json
  {
    "Message": "welcome to URL shortner"
  }
  ```

### Shorten URL

- **URL:** `/short`
- **Method:** `POST`
- **Description:** Shortens the provided URL.
- **Request Body:**
  ```json
  {
    "url": "https://example.com"
  }
  ```
- **Response:**
  ```json
  {
    "short_url": "short_url_code"
  }
  ```

### Redirect to Original URL

- **URL:** `/redirect/{short_url_code}`
- **Method:** `GET`
- **Description:** Redirects to the original URL corresponding to the short URL code.
- **Response:** HTTP 302 Redirect to the original URL.

## Example Usage

1. **Shorten a URL:**

   ```sh
   curl -X POST -d '{"url": "https://example.com"}' -H "Content-Type: application/json" http://localhost:3000/short
   ```

   **Response:**
   ```json
   {
     "short_url": "abc12345"
   }
   ```

2. **Redirect using the shortened URL:**

   Open a browser and go to `http://localhost:3000/redirect/abc12345`, and it will redirect to `https://example.com`.

## Code Overview

### Main Functions

- `welcome`: Handles the welcome endpoint.
- `shortURLHandler`: Handles the URL shortening process.
- `redirectURLHandler`: Handles redirection from a short URL to the original URL.

### URL Struct

- `ID`: The unique identifier for the URL.
- `OriginalURL`: The original URL provided by the user.
- `ShortURL`: The shortened version of the URL.
- `CreationDate`: The date and time when the URL was created.

### Helper Functions

- `generateShortURL`: Generates a short URL code using MD5 hashing.
- `createURL`: Stores the original and short URL in the database.
- `getURL`: Retrieves the original URL from the database using the short URL code.

### In-Memory Database

The URLs are stored in an in-memory map `urlDB` where the key is the short URL code and the value is the `URL` struct.

## Running Tests

Currently, there are no tests included in this project. However, you can use tools like `Postman` to manually test the API endpoints.

## Acknowledgments

- [Go Documentation](https://golang.org/doc/)
- [MD5 Package](https://pkg.go.dev/crypto/md5)
- [HTTP Package](https://pkg.go.dev/net/http)

---

Feel free to customize the content according to your repository and additional features or configurations you may have.