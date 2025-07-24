# 🚀 Self-Hosted CI/CD with Go & Webhook Integration

This project is a self-hosted CI/CD service built using **Golang** that listens for GitHub webhook events and automatically pulls and builds code from a remote repository. It is designed as a learning project and will be expanded with a full-featured UI dashboard in future development stages.

## 📚 Project Overview

The main goal is to understand and build a simple CI/CD server without relying on external tools like GitHub Actions or Jenkins. This solution is lightweight and suitable for self-hosted environments or internal deployment pipelines.

### ✨ Features

- Listens to GitHub webhooks via a RESTful endpoint
- Automatically clones/pulls latest code from a repository
- Executes shell commands to build/deploy the application
- Logs all deployment activity
- Built with Go (Golang) for high performance and concurrency

> **Note**: This is an educational project and still in development. More features such as UI monitoring, multi-branch support, and notification hooks are planned.

---

## 🔧 Tech Stack

- **Language:** Go (Golang)
- **Runtime:** Go 1.20+
- **Tools:** `git`, `bash`, `curl`
- **Deployment Target:** Any Linux-based system

---

## 🛠️ Setup Instructions

### 1. Clone the Repository

git clone https://github.com/rizzsv/SelfHosted_CICD.git
cd SelfHosted_CICD.

---

## Run the Service
go run main.go

---

## 📁 Folder Structure|
```
.
├── config/              # Konfigurasi CORS & lainnya
│   └── cors.go
├── executor/            # Eksekusi build dan run project
│   └── runner.go
├── executor logs tmp/   # Tempat log hasil eksekusi build
├── handler/             # Penanganan webhook request
│   ├── build.go
│   └── webhook.go
├── model/               # Struct model yang digunakan
│   └── build.go
├── routes/              # Routing untuk endpoint
│   └── router.go
├── storage/             # Logging utilitas
│   └── logs.go
├── tmp/                 # Tempat penyimpanan file build sementara
│   └── 20250724-070439/ # Contoh folder hasil push
├── main.go              # Entry point aplikasi
├── go.mod
└── go.sum

```
