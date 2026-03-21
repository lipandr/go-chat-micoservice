# go-chat-microservice

Chat microservice built with Go.

## Features

- HTTP server with graceful shutdown
- Structured logging with logrus
- Configuration via environment variables
- CI/CD with GitHub Actions

## Prerequisites

- Go 1.23+

## Quick Start

```bash
make build
make run
```

## Makefile Commands

| Command      | Description                |
|-------------|----------------------------|
| `make build`   | Build binary             |
| `make run`     | Build and run           |
| `make lint`    | Run linter              |
| `make lint-fix`| Auto-fix linter issues  |
| `make vet`     | Run go vet              |
| `make test`    | Run tests               |

## Configuration

| Variable      | Default   | Description        |
|---------------|-----------|---------------------|
| `SERVER_HOST` | localhost | Server host         |
| `SERVER_PORT` | 8080      | Server port         |
| `DB_HOST`     | localhost | Database host       |
| `DB_PORT`     | 5432      | Database port       |
| `DB_USERNAME` | postgres  | Database username   |
| `DB_PASSWORD` | postgres  | Database password   |
| `DB_NAME`     | chat      | Database name       |

## Endpoints

- `GET /health` - Health check

## Project Structure

```
├── cmd/main.go           # Entry point
├── internal/
│   ├── config/           # Configuration
│   ├── logger/           # Logging
│   ├── server/           # HTTP server
│   └── shutdown/         # Graceful shutdown
├── Makefile
└── .golangci.pipeline.yaml
```
