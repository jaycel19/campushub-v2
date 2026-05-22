# CampusHub

CampusHub is a modern fullstack campus communication and collaboration platform designed to simulate real-world software engineering practices and scalable backend architecture.

This project is being developed as a flagship portfolio project focused on:
- scalable backend architecture
- modular system design
- authentication & authorization
- feed ranking algorithms
- clean code practices
- production-style workflows

---

# Tech Stack

## Backend
- Go
- Gin
- PostgreSQL
- GORM
- JWT Authentication
- Goose Migrations
- Docker

## Frontend (Planned)
- React
- TypeScript
- Vite
- TailwindCSS
- TanStack Query
- Zustand

---

# 📁 Project Structure

```bash
campushub/
│
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── auth/
│   │   ├── post/
│   │   ├── comment/
│   │   ├── middleware/
│   │   ├── shared/
│   │   └── database/
│   │
│   ├── migrations/
│   ├── docker/
│   ├── .env
│   ├── go.mod
│   └── Makefile
│
├── frontend/
│   ├── src/
│   ├── public/
│   └── package.json
│
├── docker-compose.yml
└── README.md
```

---

# 🏗️ Backend Architecture

The backend follows a layered architecture approach:

```text
Handler Layer
    ↓
Service Layer
    ↓
Repository Layer
    ↓
PostgreSQL
```

---

## Handler Layer
Responsible for:
- HTTP requests
- Request validation
- Response formatting

---

## Service Layer
Responsible for:
- business logic
- authorization
- validation
- workflows
- feed ranking

---

## Repository Layer
Responsible for:
- database access
- queries
- persistence

---

# 🔐 Authentication Features

- JWT Authentication
- User Registration
- User Login
- Protected Routes
- Password Hashing using bcrypt
- Middleware-based authorization

---

# 📝 Post Features

- Create Posts
- Feed Retrieval
- Pagination
- Feed Ranking Algorithm
- UUID Primary Keys

---

# 💬 Comment Features

- Create Comments
- Retrieve Comments by Post
- Delete Own Comments
- Ownership Validation

---

# ⚡ Feed Ranking Algorithm

CampusHub currently uses a simple engagement-based ranking system:

```go
score := (likes * 2) + (comments * 3) - agePenalty
```

This helps prioritize engaging and recent posts.

Future improvements may include:
- trending algorithms
- personalized feeds
- Redis caching
- recommendation systems

---

# 🧪 Engineering Goals

This project focuses on demonstrating:
- clean architecture
- scalable backend development
- modular system design
- REST API development
- authentication systems
- database relationships
- production-ready workflows

---

# 📌 Planned Features

## Backend
- Likes System
- User Profiles
- Role-Based Access Control
- Redis Caching
- Real-time Notifications
- WebSockets
- Swagger Documentation
- Unit & Integration Tests
- Activity Logging

---

## Frontend
- Authentication Pages
- Feed UI
- Post Creation
- Comments UI
- Notifications
- Responsive Design
- Dark Mode

---

# 🐳 Running the Backend

## 1. Start PostgreSQL

```bash
docker compose up -d
```

---

## 2. Run Migrations

```bash
goose up
```

---

## 3. Start the API

```bash
go run cmd/api/main.go
```

---

# 🔑 Environment Variables

Create a `.env` file inside `/backend`

```env
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=campushub

JWT_SECRET=your_secret_key
```

---

# 📖 API Endpoints

## Authentication

| Method | Endpoint |
|---|---|
| POST | `/auth/register` |
| POST | `/auth/login` |
| GET | `/auth/me` |

---

## Posts

| Method | Endpoint |
|---|---|
| GET | `/posts/feed` |
| POST | `/posts` |

---

## Comments

| Method | Endpoint |
|---|---|
| POST | `/posts/:id/comments` |
| GET | `/posts/:id/comments` |
| DELETE | `/comments/:id` |

---

# 🎯 Project Vision

CampusHub is not intended to be a simple CRUD social media project.

The goal is to evolve it into a production-style platform that demonstrates:
- backend engineering skills
- software architecture understanding
- scalable development practices
- modern fullstack workflows

---

# 👨‍💻 Author

Developed by Jaycel Lalongisip as part of a long-term software engineering and backend architecture learning journey.