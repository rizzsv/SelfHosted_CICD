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
---
├── config/                # Konfigurasi middleware (seperti CORS)
├── executor/              # Eksekusi command dari webhook (shell execution)
├── handler/               # HTTP handler (webhook, build, dll)
├── model/                 # Struktur data / model internal
├── routes/                # Router Go
├── storage/               # Penyimpanan log
├── tmp/                   # File log hasil build
├── .bolt/                 # (Opsional) penyimpanan BoltDB
├── src/                   # Frontend UI (Vite, Tailwind, dsb)
├── index.html             # Entry point frontend
├── tailwind.config.js     # Konfigurasi Tailwind CSS
├── tsconfig*.json         # Konfigurasi TypeScript
├── vite.config.ts         # Konfigurasi Vite
├── go.mod / go.sum        # Modul Go
└── README.md              # Dokumentasi
