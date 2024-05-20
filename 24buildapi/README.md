# Course Management API

This project is a Course Management API built with Go (Golang) using the Gorilla Mux router. It provides endpoints to manage courses, including creating, updating, retrieving, and deleting courses.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Running the Server](#running-the-server)
- [API Endpoints](#api-endpoints)
- [Course Model](#course-model)
- [Contributing](#contributing)
- [License](#license)

## Overview

The Course Management API allows users to perform CRUD (Create, Read, Update, Delete) operations on courses. Each course has a unique ID, name, price, and an associated author. The API responds with JSON data.

## Features

- Retrieve all courses
- Retrieve a single course by ID
- Create a new course
- Update an existing course by ID
- Delete a course by ID


## Running the Server

Start the server using the following command:

```bash
go run main.go
```

The server will start on port 4000. You can access it at `http://localhost:4000`.

## API Endpoints

### Get All Courses

- **URL:** `/courses`
- **Method:** `GET`
- **Description:** Retrieve all courses.
- **Response:** List of courses in JSON format.

### Get One Course

- **URL:** `/course/{id}`
- **Method:** `GET`
- **Description:** Retrieve a single course by ID.
- **Response:** Course details in JSON format.

### Create One Course

- **URL:** `/course`
- **Method:** `POST`
- **Description:** Create a new course.
- **Request Body:** JSON object containing `coursename`, `price`, and `author`.
- **Response:** The created course in JSON format.

### Update One Course

- **URL:** `/course/{id}`
- **Method:** `PUT`
- **Description:** Update an existing course by ID.
- **Request Body:** JSON object containing updated `coursename`, `price`, and `author`.
- **Response:** The updated course in JSON format.

### Delete One Course

- **URL:** `/course/{id}`
- **Method:** `DELETE`
- **Description:** Delete a course by ID.
- **Response:** Success message in JSON format.

## Course Model

### Course

```go
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
```

### Author

```go
type Author struct {
	Fullname string `json:"fullnam"`
	Webiste  string `json:"website"`
}
```

### Example

```json
{
	"courseid": "1",
	"coursename": "Go Programming",
	"price": 199,
	"author": {
		"fullnam": "John Doe",
		"website": "https://johndoe.com"
	}
}
```
