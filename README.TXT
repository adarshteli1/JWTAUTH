# 🔐 JWT Authentication API

A lightweight **JWT Authentication API** built with **Go**, **Gin**, and **MongoDB**. This project demonstrates secure authentication, role-based authorization, password hashing, and a clean layered architecture suitable for backend applications.

---

## 🚀 Features

* User Registration & Login
* JWT Access & Refresh Tokens
* Password Hashing with bcrypt
* Role-Based Authorization (Admin/User)
* Protected Routes using Middleware
* MongoDB Integration
* Layered Architecture (Controller → Service → Helper)

---

## 🛠️ Tech Stack

| Technology | Purpose              |
| ---------- | -------------------- |
| Go         | Programming Language |
| Gin        | Web Framework        |
| MongoDB    | Database             |
| JWT        | Authentication       |
| bcrypt     | Password Hashing     |

---

## 📂 Project Structure

```text
JWTAUTH/
├── controllers/
├── database/
├── helpers/
├── middleware/
├── models/
├── routes/
├── service/
├── main.go
└── .env
```

---

## 🏗️ Architecture

```text
Client
   │
   ▼
Router
   │
   ▼
Controller
   │
   ▼
Service
   │
   ▼
Helpers
   │
   ▼
MongoDB
```

---

## 🌐 API Endpoints

| Method | Endpoint          | Description           |
| ------ | ----------------- | --------------------- |
| POST   | `/users/signup`   | Register a new user   |
| POST   | `/users/login`    | Authenticate user     |
| GET    | `/users`          | Get all users (Admin) |
| GET    | `/users/:user_id` | Get user by ID        |

---

## 🔒 Authentication Flow

1. User signs up with email and password.
2. Password is securely hashed using bcrypt.
3. User logs in and receives JWT Access & Refresh Tokens.
4. Middleware validates JWT for protected routes.
5. Role-based authorization restricts access to resources.

---

## ⚙️ Running the Project

```bash
git clone https://github.com/adarshteli1/JWTAUTH.git

cd JWTAUTH

go mod tidy

go run main.go
```

---

## 📈 Future Improvements

* Refresh Token Endpoint
* Logout API
* Update/Delete User
* Swagger Documentation
* Unit Testing
* Docker Support

---

## 🎯 Learning Outcomes

* JWT Authentication & Authorization
* REST API Development with Gin
* Layered Backend Architecture
* MongoDB Integration
* Password Hashing using bcrypt
* Middleware & Protected Routes

---

## 👨‍💻 Author

**Adarsh Teli**

Backend Engineering Portfolio Project built to practice secure authentication, clean architecture, and production-oriented Go backend development.
