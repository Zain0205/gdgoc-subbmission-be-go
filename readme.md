-----

# GDG On Campus – StuJam Backend

A high-performance, containerized **Go** backend designed for the GDG On Campus StuJam challenge. This project implements a robust REST API with a focus on modular architecture, utilizing **Docker Compose** for a seamless, hot-reload development environment.

## ⚡️ Tech Stack

  * **Core:** Go (Golang)
  * **Database:** MySQL 8.0
  * **Infrastructure:** Docker & Docker Compose
  * **Proxy:** Nginx
  * **Dev Tooling:** Air (Hot Reload), Vite (Frontend Integration)

## 🚀 Quick Start

Get the complete stack running in minutes.

**Prerequisites:** Docker ≥ 20.10, Git.

```bash
# 1. Clone repositories (Ensure standard naming convention)
git clone https://github.com/Zain0205/gdgoc-subbmission-be-go.git backend
git clone https://github.com/Zain0205/gdgoc-submission-fe-react.git frontend

# 2. Configure environment
cd gdgoc-subbmission-be-go/ 
cp .env.example .env

# 3. Launch Services
docker-compose up -d --build
```

The API will be available at `http://localhost:8080`.

## 🏗 Architecture & Services

The application runs as a cohesive containerized suite managed by Nginx.

[Image of containerized microservices architecture diagram]

| Service | Container Name | Technology | Internal Port | Public URL | Purpose |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Proxy** | `stujam_nginx` | Nginx | 80 | `localhost:8000` | Reverse Proxy & Routing |
| **Backend** | `stujam_backend` | Go + Air | 8080 | `localhost:8080` | REST API |
| **Frontend** | `stujam_frontend`| React + Vite | 5173 | `localhost:5173` | UI & Client Logic |
| **Database**| `stujam_db` | MySQL 8.0 | 3306 | `localhost:3306` | Persistent Storage |

## 📂 Project Structure

```text
parent-directory/
├── gdgoc-subbmission-be-go/       # Backend (this repository)
│   ├── config/
│   │   └── config.go
│   ├── controllers/
│   │   ├── achievement_controller.go
│   │   ├── auth_controller.go
│   │   ├── leaderboard_controller.go
│   │   ├── member_controller.go
│   │   ├── notification_controller.go
│   │   ├── series_controller.go
│   │   ├── submission_controller.go
│   │   ├── track_controller.go
│   │   └── user_controller.go
│   ├── database/
│   │   └── database.go
│   ├── docker/nginx/dev.conf
│   ├── dto/dto.go
│   ├── middleware/auth.go
│   ├── models/entity.go
│   ├── routes/routes.go
│   ├── uploads/
│   │   ├── avatars/
│   │   └── badges/
│   ├── utils/
│   │   ├── file_utils.go
│   │   ├── jwt.go
│   │   └── response.go
│   ├── validation/validator.go
│   ├── docker-compose.yml
│   ├── Dockerfile.dev
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── openapi.yaml
│   ├── .env.example
│   └── readme.md
└── gdgoc-submission-fe-react/     # Frontend repository
    ├── src/
    │   ├── components/
    │   ├── pages/
    │   ├── hooks/
    │   └── services/
    ├── Dockerfile.dev
    └── package.json
```

## 🛠 Development Workflow

### Common Commands

| Action | Command |
| :--- | :--- |
| **Start Stack** | `docker-compose up -d` |
| **Follow Logs** | `docker-compose logs -f [service_name]` |
| **Rebuild** | `docker-compose up -d --build` |
| **Stop** | `docker-compose down` |
| **Reset DB** | `docker-compose down -v` |

### Troubleshooting

  * **Port Conflicts:** Ensure ports `8000`, `8080`, `5173`, and `3306` are free, or modify `docker-compose.yml`.
  * **Permissions:** If you encounter write errors on Linux: `sudo chown -R $USER:$USER .`
  * **Hot-Reload:** If changes aren't reflecting, ensure `.air.toml` is correctly configured in the root.

## 🤝 Contributing

Contributions are welcome\! Please follow these steps:

1.  Fork the repository.
2.  Create a feature branch (`git checkout -b feature/amazing-feature`).
3.  Commit your changes.
4.  Open a Pull Request.

## 📜 Acknowledgements

Developers for the **GDG On Campus** submission app.

  * **Team:** [@Zain0205](https://github.com/Zain0205), [@rhankbrguw](https://github.com/rhankbrguw), [@sepUnch](https://github.com/sepUnch) 
  * **Special Thanks:** The GDG Organizing Team for the platform and support.

---
