# Library Manager App Backend

A comprehensive Golang backend service that provides two main functionalities:

1. **RESTful API for Book Management** - Full CRUD operations for managing books
2. **URL Cleanup and Redirection Service** - Process URLs for canonical formatting and redirection

Built with Go, featuring database integration, comprehensive validation, detailed logging, and Swagger documentation.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Database Setup](#database-setup)
  - [Running the application](#running-the-application)
- [API Documentation](#api-documentation)
- [License](#license)

## Features

### Book Management API

- **Full CRUD Operations**: Create, Read, Update, Delete books
- **Database Integration**: Persistent data storage with any SQL database
- **Input Validation**: Comprehensive backend validation for all fields
- **Error Handling**: Detailed error responses with appropriate HTTP status codes
- **Structured Logging**: Request/response logging with timestamps

### URL Processing Service

- **Canonical URL Processing**: Remove query parameters and trailing slashes
- **Redirection Processing**: Convert to lowercase and ensure www.byfood.com domain
- **Combined Operations**: Support for both operations simultaneously
- **Input Validation**: Rigorous validation with detailed error messages

### General Features

- **Swagger Documentation**: Interactive API documentation
- **Comprehensive Testing**: Unit tests with edge case coverage
- **Middleware**: Request logging, error handling, and recovery

## Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin-Gonic
- **Database**: PostgreSQL
- **ORM**: GORM
- **Documentation**: Swagger
- **Testing**: Go testing framework with testify
- **Logging**: Logrus structured logging
- **Configuration**: Viper for environment management
- **Migration**: dbmate

## Project Structure

```
backend/
├── cmd/
│   └── app/
│       └── main.go           # Application entry point
├── config/                   # Configuration management
├── db/
│   └── migrations/           # Database migration
├── docs/                     # Swagger documentation
├── generated/
│   └── mockgen/              # Generated mockgen mock
├── internal/
│   ├── book/                 # Book domain module
│   │   ├── entity/           # Business entity
│   │   ├── handler/          # HTTP handlers
│   │   │   ├── dto/          # Data Transfer Objects for request and response
│   │   ├── mapper/           # Entity-DTO mapping
│   │   ├── repo/             # Data access layer
│   │   │   └── model/        # Database models
│   │   └── usecase/          # Business logic layer
│   └── url/                  # URL processing domain module
│       └── ...               # Same pattern as book/ domain without repo
├── pkg/                      # Reusable, self-contained internal packages
├── .env.example              # Environment variables template need to rename to .env
├── go.mod                    # Go module dependencies
└── go.sum                    # Go module checksums
```

## Setup Instructions

### Prerequisites

- **Go 1.21+** installed

### Installation

1. Make sure current directory is `/backend`

2. Create a `.env` file in the root directory based on the `.env.sample`

3. Install dependencies:

```bash
go mod download
go mod tidy
```

4. Install development tools:

```bash
# Install swagger generator
go install github.com/swaggo/swag/cmd/swag@latest

# Install mockgen generator
go install go.uber.org/mock/mockgen@latest

# Install dbmate
go install github.com/amacneil/dbmate/v2/cmd/dbmate@latest

# Add GORM library to your project
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### Database setup

1. Set the value of PostgreSQL related in the `.env` file.

2. Initiate the database:

```bash
docker-compose up
```

3. Migrate the database:

```bash
# Run migrations
dbmate up
```

### Running the Application

1. Run the development server:

```bash
go run cmd/app/main.go
```

2. The application will be available at:

- **API**: `http://localhost:8080`
- **Swagger UI**: `http://localhost:8080/swagger/index.html`

## API Documentation

### Swagger UI

Access the interactive API documentation at: `http://localhost:8080/swagger/index.html`

### Generating Documentation

```bash
swag init -g cmd/app/main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
