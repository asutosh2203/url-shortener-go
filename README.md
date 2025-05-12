# ⚡ URL Shortener in Go

A lightweight URL shortener service built with Go, Gin, and Redis. This project allows users to convert long URLs into short, shareable links — complete with persistent storage via Redis.

---

## 🚀 Features

- 🔗 Shortens long URLs into 5-character codes
- 💾 Persists URL mappings using Redis
- 🔁 Supports redirection from short URLs to original URLs
- ⏳ Allows URL expiration by specifying a TTL (Time To Live) in hours
- 🧪 Easy to test using tools like `curl` or `httpie`
- 🛡️ Basic IP-based rate limiting (max 5 requests per minute per IP)

---

## 🛠 Tech Stack

- **Go** (Golang)
- **Gin** (HTTP web framework)
- **Redis** (for key-value storage)

---

## 📦 Project Structure

url-shortener-go/  
│  
├── main.go              --> Entry point  
├── handlers/  
│   ├── shorten.go       --> Logic for generating and returning shortened URLs  
│   └── redirect.go      --> Logic for handling redirection from short to long URL  
├── storage/  
│   └── redis.go         --> Redis setup and get/set helpers  
├── middleware/  
│   └── rate_limiter.go  --> IP based rate limiter using Redis  
├── utils/  
│   └── utils.go         --> Utility helpers  

---

## ⚙️ Getting Started

### 1. Clone the repo

``` cmd
git clone https://github.com/asutosh2203/url-shortener-go.git
cd url-shortener-go
```

### 2. Install dependencies

Make sure Go is installed. Then:

```cmd
go mod tidy
```

### 3. Start Redis (if using WSL or Docker)

If using WSL with Redis installed

```cmd
redis-server
```

### 4. Run the server

```cmd
go run main.go
```

---

## 🔁 Example Usage
### Shorten a URL

```cmd
curl -X POST http://localhost:8080/shorten \ 
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com", "ttl": 24}'
```

In this example, the shortened URL will expire after 24 hours. If no TTL is provided, the URL will remain valid indefinitely.

### Response:

```json
{
  "message": "URL shortened successfully",
  "shortUrl": "localhost:8080/abc12"
}
```

### Visit the short URL

```cmd
curl -L http://localhost:8080/abc12
```

Or open it in your browser — it redirects to the original long URL.

---

### 🛡️ Rate Limiting

To prevent abuse, this app uses a basic Redis-backed rate limiter. Each IP address can make up to 5 requests per minute. After that, you'll receive a `429 Too Many Requests` response.

---

## 📌 Notes

    - Redis keys are now stored with an optional expiration time (TTL). If no TTL is provided, the shortened URL mapping persists indefinitely.

    - This is a beginner-friendly implementation; advanced features like analytics and user authentication may be added later.
