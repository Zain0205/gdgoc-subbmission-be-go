# GDG OC Submission - Go Backend 🚀

A containerized Go backend service for the GDG OC Submission application, with Docker Compose orchestration for streamlined development.

## Prerequisites

- [Docker](https://www.docker.com/get-started) 🐳
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Git](https://git-scm.com/)

## Project Structure

Ensure the following directory structure for proper Docker Compose operation:

```
parent-directory/
├── gdgoc-subbmission-be-go/    # Backend repository
│   ├── docker-compose.yml
│   ├── Dockerfile.dev
│   ├── docker/
│   │   └── nginx/
│   │       └── dev.conf
│   ├── .env.example
│   ├── go.mod
│   └── ...
└── gdgoc-submission-fe-react/  # Frontend repository
    ├── Dockerfile.dev
    ├── package.json
    └── ...
```

## Getting Started

### 1. Clone Repositories

Clone both backend and frontend repositories as siblings:

```bash
# Clone backend
git clone https://github.com/Zain0205/gdgoc-subbmission-be-go.git

# Clone frontend (sibling directory)
git clone https://github.com/Zain0205/gdgoc-submission-fe-react.git
```

### 2. Configure Environment

Navigate to the backend directory and create your environment file:

```bash
cd gdgoc-subbmission-be-go
cp .env.example .env
```

Update `.env` with your configuration. Required variables:

```env
DB_DATABASE=your_database_name
DB_USERNAME=your_db_user
DB_PASSWORD=your_db_password
DB_ROOT_PASSWORD=your_root_password
```

### 3. Initialize Dependencies (Optional)

If `go.sum` is missing or dependencies have changed:

```bash
go mod tidy
```

## Running the Application

### Start Services

From the `gdgoc-subbmission-be-go` directory:

```bash
docker-compose up -d --build
```

- `--build`: Rebuilds images if changes are detected
- `-d`: Runs containers in detached mode

### Stop Services

```bash
docker-compose down
```

### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
```

## Access Points

Once running, access the application at:

- **Application (via Nginx)**: http://localhost:8000 🌐
- **Backend Direct**: http://localhost:8080
- **Frontend Direct**: http://localhost:5173
- **MySQL Database**: `localhost:3307` (internal port: 3306) 💾

## Service Architecture

The application consists of four services:

1. **Backend** (`stujam_backend`): Go service with Air for hot reloading on port 8080
2. **Frontend** (`stujam_frontend`): React with Vite dev server on port 5173
3. **Database** (`stujam_db`): MySQL 8.0 with persistent volume storage
4. **Nginx** (`stujam_nginx`): Reverse proxy routing traffic on port 8000

## Development Workflow

### Hot Reloading ⚡

- **Backend**: Uses [Air](https://github.com/air-verse/air) for automatic recompilation on `.go` file changes
- **Frontend**: Vite dev server provides Hot Module Replacement (HMR) for instant React updates

Changes to source files are automatically detected and reflected without manual restarts.

### Database Persistence

MySQL data is persisted in a Docker volume (`mysql_data`). To reset the database:

```bash
docker-compose down -v  # Removes volumes
docker-compose up -d --build
```

## Troubleshooting

### Port Conflicts

If ports 8000, 8080, 5173, or 3307 are already in use, update the port mappings in `docker-compose.yml`:

```yaml
ports:
  - "NEW_PORT:CONTAINER_PORT"
```

### Permission Issues

If you encounter permission errors with volumes:

```bash
sudo chown -R $USER:$USER .
```

### Rebuild Containers

To force a complete rebuild:

```bash
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

---

For issues or contributions, please refer to the project's [issue tracker](https://github.com/Zain0205/gdgoc-subbmission-be-go/issues). 📝
