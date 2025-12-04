```markdown
# GDG On Campus Submission – Go Backend

A production-grade, containerized Go backend with hot-reload development environment, built for the GDG On Campus StuJam challenge.

## Prerequisites

- Docker ≥ 20.10
- Docker Compose ≥ v2.0
- Git

## Project Structure
```

parent-directory/
├── gdgoc-subbmission-be-go/ # Backend (this repository)
│ ├── config/
│ │ └── config.go
│ ├── controllers/
│ │ ├── achievement_controller.go
│ │ ├── auth_controller.go
│ │ ├── leaderboard_controller.go
│ │ ├── member_controller.go
│ │ ├── notification_controller.go
│ │ ├── series_controller.go
│ │ ├── submission_controller.go
│ │ ├── track_controller.go
│ │ └── user_controller.go
│ ├── database/
│ │ └── database.go
│ ├── docker/nginx/dev.conf
│ ├── dto/dto.go
│ ├── middleware/auth.go
│ ├── models/entity.go
│ ├── routes/routes.go
│ ├── uploads/
│ │ ├── avatars/
│ │ └── badges/
│ ├── utils/
│ │ ├── file_utils.go
│ │ ├── jwt.go
│ │ └── response.go
│ ├── validation/validator.go
│ ├── docker-compose.yml
│ ├── Dockerfile.dev
│ ├── go.mod
│ ├── go.sum
│ ├── main.go
│ ├── openapi.yaml
│ ├── .env.example
│ └── readme.md
└── gdgoc-submission-fe-react/ # Frontend repository
├── src/
│ ├── components/
│ ├── pages/
│ ├── hooks/
│ └── services/
├── Dockerfile.dev
└── package.json

````

## Quick Start

```bash
git clone https://github.com/Zain0205/gdgoc-subbmission-be-go.git
git clone https://github.com/Zain0205/gdgoc-submission-fe-react.git

cd gdgoc-subbmission-be-go
cp .env.example .env
docker-compose up -d --build
````

## Service Endpoints

| Service  | URL                   | Purpose         |
| -------- | --------------------- | --------------- |
| Nginx    | http://localhost:8000 | Reverse proxy   |
| Backend  | http://localhost:8080 | API             |
| Frontend | http://localhost:5173 | Vite dev server |
| MySQL    | localhost:3306        | Database        |

## Architecture

| Service  | Container       | Technology   | Port |
| -------- | --------------- | ------------ | ---- |
| Backend  | stujam_backend  | Go + Air     | 8080 |
| Frontend | stujam_frontend | React + Vite | 5173 |
| Database | stujam_db       | MySQL 8.0    | 3306 |
| Proxy    | stujam_nginx    | Nginx        | 80   |

## Development Features

- Backend hot-reload via Air (v1.62.0 – pinned for stability)
- Frontend instant updates via Vite HMR
- Full local stack including MySQL and Nginx reverse proxy

## Useful Commands

```bash
docker-compose logs -f [service]      # Follow logs
docker-compose down                   # Stop services
docker-compose down -v                # Remove volumes (reset DB)
docker-compose up -d --build          # Rebuild & start
docker-compose build --no-cache       # Force rebuild
```

## Troubleshooting

- Port already in use → Modify `docker-compose.yml` port mappings
- Permission errors → `sudo chown -R $USER:$USER .`
- Hot-reload issues → Check backend logs and ensure `.air.toml` is present

## Production Recommendations

- Multi-stage Dockerfile for minimal image size
- Serve static React build through Nginx
- Use secret management (Vault, AWS Secrets Manager, etc.)
- Enable database connection pooling

## Acknowledgements

We extend our sincere gratitude to:

- [@rhankbrguw](https://github.com/rhankbrguw)
- [@sepUnch](https://github.com/sepUnch)

For their significant contributes, code reviews.

Special thanks to the **Google Developer Groups on Campus** organizing team for providing this invaluable platform and learning opportunity.

## Contributing

Issues and pull requests are welcome.  
Please use the GitHub issue tracker and follow standard contribution practices.

---
