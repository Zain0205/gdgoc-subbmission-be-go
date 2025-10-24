# 🚀 GDG OC Submission – Go Backend

A containerized Go backend service with Docker-based development environment featuring hot-reload capabilities.

## 📋 Prerequisites

- [Docker](https://www.docker.com/get-started) (20.10+)
- [Docker Compose](https://docs.docker.com/compose/install/) (v2.0+)
- [Git](https://git-scm.com/)

## 📁 Project Structure

```
parent-directory/
├── gdgoc-subbmission-be-go/      # Backend (this repository)
│   ├── docker-compose.yml
│   ├── Dockerfile.dev
│   ├── docker/nginx/dev.conf
│   └── .env.example
└── gdgoc-submission-fe-react/    # Frontend repository
    ├── Dockerfile.dev
    └── package.json
```

## ⚡ Quick Start

### 1. Clone Repositories

```bash
git clone https://github.com/Zain0205/gdgoc-subbmission-be-go.git
git clone https://github.com/Zain0205/gdgoc-submission-fe-react.git
```

### 2. Environment Configuration

```bash
cd gdgoc-subbmission-be-go
cp .env.example .env
```

Required environment variables:

```bash
# Application
APP_ENV=development
APP_PORT=8080

# Database
DB_HOST=db
DB_PORT=3306
DB_DATABASE=stujam_db
DB_USERNAME=stujam_user
DB_PASSWORD=stujam_password
DB_ROOT_PASSWORD=your_root_password_here

# Security
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# CORS
ALLOWED_ORIGINS=http://localhost:8000,http://localhost:5173
```

### 3. Start Services

```bash
docker-compose up -d --build
```

## 🌐 Service Endpoints

| Service  | URL                         | Purpose                   |
|----------|-----------------------------|---------------------------|
| Nginx    | `http://localhost:8000`     | Reverse proxy entry point |
| Backend  | `http://localhost:8080`     | Direct API access         |
| Frontend | `http://localhost:5173`     | Development server        |
| MySQL    | `localhost:3306`            | Database connection       |

## 🏗️ Architecture

| Service  | Container          | Technology   | Port |
|----------|--------------------|--------------|------|
| Backend  | `stujam_backend`   | Go + Air     | 8080 |
| Frontend | `stujam_frontend`  | React + Vite | 5173 |
| Database | `stujam_db`        | MySQL 8.0    | 3306 |
| Proxy    | `stujam_nginx`     | Nginx        | 80   |

## 💻 Development

### Hot Reload

- **Backend:** Air (v1.62.0) automatically reloads on `.go` file changes
- **Frontend:** Vite HMR for instant updates

> **⚠️ Note:** Air version is pinned to v1.62.0 in `Dockerfile.dev` due to compatibility issues with latest versions. Do not update without testing.

### Common Commands

```bash
# View logs
docker-compose logs -f [service-name]

# Stop services
docker-compose down

# Reset database
docker-compose down -v
docker-compose up -d --build

# Rebuild containers
docker-compose build --no-cache
```

## 🔧 Troubleshooting

**Port conflicts:** Modify port mappings in `docker-compose.yml`

**Permission errors:**
```bash
sudo chown -R $USER:$USER .
```

**Dependency updates:**
```bash
go mod tidy
```

**Air hot-reload not working:**
- Ensure Air v1.62.0 is installed (check `Dockerfile.dev`)
- Verify `.air.toml` configuration exists
- Check container logs: `docker-compose logs -f backend`

## 🚢 Production Deployment

This configuration is for development only. Production deployment requires:

1. Multi-stage Docker build for optimized Go binary
2. Static React build served via Nginx
3. Secure environment variable management
4. Database connection pooling and optimization

## 🤝 Contributing

Submit issues and pull requests via the [issue tracker](https://github.com/Zain0205/gdgoc-subbmission-be-go/issues).

---
