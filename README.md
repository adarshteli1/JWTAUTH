# 🔐 JWT Authentication API

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Framework-Gin-00A86B?style=for-the-badge)
![MongoDB](https://img.shields.io/badge/Database-MongoDB-47A248?style=for-the-badge&logo=mongodb)
![JWT](https://img.shields.io/badge/Auth-JWT-orange?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Completed-brightgreen?style=for-the-badge)

---

A secure **RESTful Authentication API** built using **Go**, **Gin**, **MongoDB**, and **JWT**. The project demonstrates authentication, authorization, password hashing, protected routes, and a clean layered backend architecture.

---

# 📚 Table of Contents

- Overview
- Features
- Tech Stack
- Project Structure
- Architecture
- API Endpoints
- Running the Project
- Future Improvements

---

# 🚀 Features

- JWT Authentication
- User Signup & Login
- Access & Refresh Tokens
- Password Hashing (bcrypt)
- Role-Based Authorization
- JWT Middleware
- MongoDB Integration
- Layered Architecture

---

# 🛠 Tech Stack

| Technology | Purpose |
|------------|---------|
| Go | Programming Language |
| Gin | Web Framework |
| MongoDB | Database |
| JWT | Authentication |
| bcrypt | Password Hashing |

---

# 📂 Project Structure

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

# 🏗 Architecture

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

# 🌐 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/users/signup` | Register User |
| POST | `/users/login` | Login User |
| GET | `/users` | Get All Users (Admin) |
| GET | `/users/:user_id` | Get User |

---

# ⚙️ Running

```bash
git clone https://github.com/adarshteli1/JWTAUTH.git

cd JWTAUTH

go mod tidy

go run main.go
```

---

# 📈 Future Improvements

- Refresh Token Endpoint
- Logout API
- Swagger Documentation
- Docker
- Unit Testing

---

# 👨‍💻 Author

**Adarsh Teli**
