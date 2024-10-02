# go-template
{{ cookiecutter.project_description }}

# how to run
1) install swagger cli 
2) `swag init`
3) `go mod tidy`
4) local: `GIN_MODE=release go run main.go`
5) or docker compose: `docker compose up -d`


## swagger setup
install go swagger cli: `go install github.com/swaggo/swag/cmd/swag@latest`
run `swag init` for document generation.

## project structure
```bash
.
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal
│   ├── handlers
│   │   └── {{cookiecutter.project_slug}}.go
│   ├── middlewares
│   │   ├── auth.go
│   │   └── handlers.go
│   ├── models
│   │   └── {{cookiecutter.project_slug}}.go
│   ├── routes
│   │   └── {{cookiecutter.project_slug}}.go
│   ├── database.go
│   ├── log.go
│   └── settings.go
├── Dockerfile
├── LICENSE
├── README.md
├── docker-compose.yaml
├── go.mod
├── go.sum
└── main.go
```