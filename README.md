# Go-Rest-API

## Introduction
Welcome to the Go-Rest-API repository! This project is a RESTful API built with Go (Golang)'s Gin Web Framework, designed to provide a simple and efficient way to handle web requests.

## Features
- CRUD operations: Create, Read, Update, and Delete resources
- JSON request and response bodies
- Middleware support for logging and authentication
- Test requests for reliability and maintainability
- JSON Web Tokens (JWT) to handle authentication and authorization


## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go 1.x.x

### Installing
Clone the repository to your local machine:
```bash
git clone https://github.com/HelloHvbert/go-rest-api.git
```
Navigate to the project directory:
```bash
cd go-rest-api/REST_API
```
Install dependencies:
```bash
go mod tidy
```
Start the server:
```bash
go run main.go
```

### Running with Docker
Navigate to the project directory:
```bash
cd go-rest-api/REST_API
```
Build the Docker image:
```bash
docker build -t go-rest-api .
```
Run the Docker container:
```bash
docker run -d -p 8080:8080 go-rest-api
```

The application will now be accessible at http://localhost:8080.
