# Ticket Management System

A backend Ticket Management System built using **Golang**, **Gin**, **GORM**, **SQLite**, and **JWT Authentication**.

## Features

- User Registration
- User Login
- JWT Authentication
- Password Hashing (bcrypt)
- Create Ticket
- Get All User Tickets
- Get Ticket By ID
- Update Ticket Status
- Ownership-based Authorization
- Status Transition Validation
- SQLite Database
- REST APIs

---

## Tech Stack

- Golang
- Gin Framework
- GORM
- SQLite
- JWT
- bcrypt

---

## API Endpoints

| Method | Endpoint |
|---------|----------|
| GET | /health |
| POST | /auth/register |
| POST | /auth/login |
| POST | /tickets |
| GET | /tickets |
| GET | /tickets/:id |
| PATCH | /tickets/:id/status |

---

## Status Flow

```text
open -> in_progress -> closed

closed tickets cannot be reopened
```

---

## Run Locally

```bash
go mod tidy
go run ./cmd
```

Server starts on:

```
http://localhost:8080
```

Health API:

```
GET /health
```

---

## Author

**Ritik Pareek**

GitHub:
https://github.com/99RitikPareek