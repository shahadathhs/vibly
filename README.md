# 🌐 **Vibly**

### ⚡ Real-Engineering-Focused Local Live Streaming Platform

**Vibly** is a high-performance live streaming platform built with **Go** and **FFmpeg**. It is designed to demonstrate backend systems thinking, media pipeline understanding, and scalability concepts using a monolithic but scalable architecture.

📌 **Swagger Docs:**
👉 `/docs/index.html`

## ✨ **Core Features**

- 🎥 **Live Stream Ingestion** – Accept live streams via RTMP (OBS compatible)
- 🎞️ **Multi-Quality Streaming** – Adaptive Bitrate (ABR) transcoding into 1080p, 720p, and 480p
- ⏺️ **Stream Recording** – Automatic local recording of all live streams
- 🖼️ **Thumbnail Generation** – Automated extraction of stream thumbnails
- 🏗️ **Worker Isolation** – Isolated FFmpeg processes for each stream to ensure system stability
- 🔐 **Authentication** – Secure JWT-based login & registration for stream management

## 🧰 **Tech Stack**

- 🏎️ **Go (1.25+)**
- 🎥 **FFmpeg** (External binary)
- 📦 **Docker & Docker Compose**
- 🔁 **Air (Live Reload)**
- 🛠️ **Makefile** for workflow automation
- ⚙️ **Lefthook** for Git hooks

## 🚀 Getting Started

### 1️⃣ **Clone the Repository**

```bash
git clone https://github.com/shahadathhs/vibly.git
cd vibly
```

### 2️⃣ **Environment Setup**

Create `.env` file:

```bash
PORT=8080
GO_ENV=development
JWT_SECRET=your_super_secret_key_here
```

## 🐳 Run Using Docker (Recommended)

### ▶️ Development Mode (with Live Reload)

```bash
make up-dev
```

👉 Runs at: **[http://localhost:8081](http://localhost:8081)**

### ▶️ Production Mode

```bash
make up
```

👉 Runs at: **[http://localhost:8080](http://localhost:8080)**

## 🔐 **Authentication Endpoints**

### ➕ Register:

**POST** `/api/auth/register`
Body:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

### 🔑 Login:

**POST** `/api/auth/login`

---

## 🏗️ **Project Structure**

```
vibly/
├── app/
│   ├── handlers/       # HTTP handlers
│   ├── middleware/     # Auth, logger, etc.
│   ├── models/         # Data models
│   └── store/          # JSON-based storage
├── pkg/
│   ├── config/         # Configuration loading
│   └── utils/          # Helpers
├── data/               # Local media and data storage
├── docs/               # Swagger documentation
├── scripts/            # Helper scripts
├── Dockerfile
├── compose.yaml
├── Makefile
└── main.go
```

## 🛠️ **Development Commands**

- 📘 `make help` – See all commands
- ▶️ `make run` – Run locally
- 🔨 `make build-local` – Build binary
- ✨ `make fmt` – Format code
- 🔍 `make vet` – Static analysis
- 🧹 `make tidy` – Cleanup modules
- 🔐 `make g-jwt` – Generate JWT secret
