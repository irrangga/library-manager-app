# Library Manager App Frontend

This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app). A React/Next.js and TypeScript frontend application to manage a book collection with full CRUD functionality and user-friendly UI.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the application](#running-the-application)
- [License](#license)

## Tech Stack

- **Frontend Framework**: Next.js 14+ with App Router
- **Language**: TypeScript
- **Styling**: Tailwind CSS / CSS Modules
- **State Management**: React Context API
- **Form Handling**: Controlled components with validation
- **HTTP Client**: Fetch API / Axios
- **Development**: ESLint, Prettier

## Features

- **Book Dashboard**: View all books in a comprehensive list
- **CRUD Operations**: Add, edit, view details, and delete books
- **Modal Forms**: Enhanced UX with modal dialogs for form submissions
- **Client-side Validation**: Real-time validation with visual feedback
- **Dynamic Routing**: Individual book detail pages with dynamic URLs
- **State Management**: Context API for global state management
- **Error Handling**: User-friendly error messages for network and form errors
- **Responsive Design**: Mobile-friendly interface
- **TypeScript**: Full type safety throughout the application

## Project Structure

```
frontend/
├── public/                   # Static assets
├── src/
│   ├── app/                  # Next.js App Router
│   │   ├── api/              # API routes (Next.js API endpoints)
│   │   ├── books/            # Books listing page (/books)
│   │   │   └── [id]/         # Dynamic book detail page (/books/[id])
│   │   ├── layout.tsx        # Root layout component
│   │   ├── page.tsx          # Home page (/)
│   │   └── globals.css       # Global styles
│   ├── components/           # React components
│   │   ├── shared/           # Reusable shared components
│   │   └── ui/               # UI-specific components (shadcn/ui)
│   └── lib/                  # Utility libraries
│       ├── api/              # API client functions
│       ├── config/           # Configuration files
│       ├── context/          # React Context providers
│       ├── definitions/      # TypeScript type definitions
│       └── utils/            # Utility functions
├── .env.sample               # Environment variables template need to rename to .env
└── package.json
```

### Route Structure Explanation

The project uses Next.js App Router with file-based routing:

- `/` - Home page (`src/app/page.tsx`)
- `/books` - Books listing page (`src/app/books/page.tsx`)
- `/books/[id]` - Individual book detail page (`src/app/books/[id]/page.tsx`)
- `/api/books` - API endpoints for book operations (`src/app/api/books/`)

Each folder in the `app` directory represents a route segment, with `page.tsx` files defining the UI for that route.

## Setup Instructions

### Prerequisites

- Node.js (v18 or higher recommended)
- pnpm package manager (see this docs [why-pnpm-package-manager](https://nextjs.org/learn/dashboard-app/getting-started))

### Installation

1. Make sure current directory is `/frontend`.

2. Create a `.env` file in the root directory based on the `.env.sample`.

3. Install dependencies:

```bash
pnpm install
```

### Running the application

1. Run the development server:

```bash
pnpm dev
```

2. Open your browser and visit `http://localhost:3000` to view the app.

## License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
