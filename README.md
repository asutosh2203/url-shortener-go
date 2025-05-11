# ⚡ URL Shortener in Go

A lightweight URL shortener service built with Go, Gin, and Redis. This project allows users to convert long URLs into short, shareable links — complete with persistent storage via Redis.

---

## 🚀 Features

- 🔗 Shortens long URLs into 5-character codes
- 💾 Persists URL mappings using Redis
- 🔁 Supports redirection from short URLs to original URLs
- 🧪 Easy to test using tools like `curl` or `httpie`

---

## 🛠 Tech Stack

- **Go** (Golang)
- **Gin** (HTTP web framework)
- **Redis** (for key-value storage)

---

## 📦 Project Structure

url-shortener-go/
│
├── main.go # Entry point
├── handlers/
│ ├── shorten.go # Logic for generating and returning shortened URLs
│ └── redirect.go # Logic for handling redirection from short to long URL
├── storage/
│ └── redis.go # Redis setup and get/set helpers


---

## ⚙️ Getting Started

### 1. Clone the repo

`git clone https://github.com/asutosh2203/url-shortener-go.git`
`cd url-shortener-go`

### 2. Install dependencies

Make sure Go is installed. Then:

`go mod tidy`

### 3. Start Redis (if using WSL or Docker)

If using WSL with Redis installed
`redis-server`

### 4. Run the server

`go run main.go`

## 🔁 Example Usage
### Shorten a URL

`curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com"}'`

### Response:

`{
  "message": "URL shortened successfully",
  "shortUrl": "localhost:8080/abc12"
}`

### Visit the short URL

`curl -L http://localhost:8080/abc12`

Or open it in your browser — it redirects to the original long URL.

## 📌 Notes

    - Redis keys are stored without expiration, ensuring persistent mapping unless manually deleted.

    - The current version is focused on core functionality; advanced features like rate limiting and analytics will come later.
