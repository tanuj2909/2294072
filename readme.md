# 🔗 URL Shortener with Redis and Fiber

A simple and fast URL shortener built with Go, using Fiber web framework and Redis as the backend for storing shortened URLs. Includes custom logging middleware.

---

## 🚀 Features

- Shorten any URL with a unique key
- Store and retrieve data from Redis
- Custom middleware for logging request details and response time
- `.env`-based configuration

---

## 📦 Requirements

- Go 1.18+
- Docker (for Redis)
- Redis server (locally or in Docker)

---

## 🛠️ Setup
To start DB:
docker run -d \
  --name db \
  -p 6379:6379 \
  -v "$(pwd)/.data":/data \
  db

To run server: go run ./api/main.go
