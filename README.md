# Library Manager App

Full-fledged web-based CRUD application for managing a small library of books. It allows users to create, read, update, and delete book records through an intuitive interface along with URL processing services.

## Architecture Overview

- **Frontend**: React/Next.js with TypeScript - Book management dashboard
- **Backend**: Golang RESTful API - Book CRUD operations + URL processing service

## Key Features

### Book Management

- **Full CRUD Operations**: Create, read, update, and delete books
- **Interactive Dashboard**: Modern UI for managing book collections
- **Form Validation**: Client-side and server-side validation
- **Modal Dialogs**: Enhanced user experience with modal forms

### URL Processing Service

- **URL Canonicalization**: Clean URLs by removing query parameters and trailing slashes
- **URL Redirection**: Convert URLs to lowercase and ensure proper domain formatting
- **Flexible Operations**: Support for individual or combined processing operations

## Quick Start

### Prerequisites

- **Node.js** 18+ (for frontend)
- **Go** 1.21+ (for backend)
- **PostgreSQL** 12+ (for database)

1. **Clone the repository**

   ```bash
   git clone https://github.com/irrangga/library-manager-app.git
   cd library-manager-app
   ```

2. **Start Backend**

   ```bash
   cd backend
   # See backend/README.md for detailed setup
   ```

3. **Start Frontend**
   ```bash
   cd frontend
   # See frontend/README.md for detailed setup
   ```

## Documentation

- **[Frontend Documentation](./frontend/README.md)** - React/Next.js setup and components
- **[Backend Documentation](./backend/README.md)** - Golang API, database setup, and endpoints

**For detailed setup instructions, API documentation, and development guidelines, see the individual README files in each directory.**
