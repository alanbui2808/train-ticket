# 🚄 E-commerce Train Ticket Booking API (Go)

```
======================================================
|        🛠️🛠️🛠️ PROJECT IS STILL ONGOING 🛠️🛠️🛠️        |
======================================================
```

Welcome to my backend API project for an e-commerce platform that allows users to book train tickets. This project is built with **Go** and follows modern backend development practices, with a strong focus on scalability, maintainability, and clean architecture.

By the end of this project, we aim to simulate a real-world train ticket booking service, demonstrating that our backend APIs can handle high volumes of requests efficiently and reliably.
---

## 🛠️ Technologies Used

- 🧠 **Go (Golang)** — Core language for backend development  
- 🌐 **Gin** — HTTP web framework for building APIs  
- 🗃️ **GORM** — ORM for interacting with MySQL  
- 🐬 **MySQL** — Relational database  
- 🐳 **Docker** — Containerization for development & deployment
- 🗄️ **Redis** — In-memory data store for caching and sessions  
- 📩 **Kafka** — Distributed event streaming for asynchronous communication    
- 🔐 **JWT** — Token-based user authentication  
- 🛡️ **RBAC** — Role-based access control  
- ⚙️ **Middleware** — Logging, error handling, validation  
- ✅ **Testing** — Unit and integration tests

---

## ✨ Features

- 🔐 **JWT-based User Authentication**
- 🧑‍💼 **Role-based Access Control** (e.g., user vs. admin)
- 🛤️ **Train Route & Schedule Management**
- 🎟️ **Ticket Booking & Order Handling**
- 🧾 **Admin Inventory Management**
- 📦 **Database Migrations & Structured Schema**

---

## 📁 Project Goals

> Build a clean, modular backend system with real-world features using Go and best practices in system design and testing.

---

## 🐳 Project Setup (Docker + Docker Compose)

This project uses **Docker** to containerize the backend and **Docker Compose** to manage MySQL and Redis for local development.

### 📦 Requirements

- Docker
- Docker Compose

### 🧱 Project Structure

```
.
├── config/
│   └── local.yaml          # App configuration
├── migrations/
│   └── shopdevgo.sql       # MySQL schema (executed at first container run)
├── docker-compose.yaml      # MySQL + Redis
├── Dockerfile              # Go backend build
```

### 🐬 Start MySQL & Redis

```bash
docker-compose up -d
```

- Starts `mysql_container` and `redis_container`
- Loads schema from `shopdevgo.sql` if volume is fresh
  
![Screenshot 2025-05-01 at 5 58 53 PM](https://github.com/user-attachments/assets/17fd0523-2eab-4ea1-872a-7ac6220190ef)

### 🏗️ Build & Run the App

```bash
docker build -t crm.shopdev.com .
docker run --rm --network=trainticket_default crm.shopdev.com
```

> Replace `trainticket_default` if your Docker Compose uses a different network name.

### ⚙️ Sample `config/local.yaml`

```yaml
mysql:
  host: mysql
  port: 3306
  dbname: shopdevgo

redis:
  host: redis
  port: 6379
```

### ♻️ Reset MySQL Schema

```bash
docker-compose down -v
docker-compose up -d
```

---

Feel free to explore the code!

