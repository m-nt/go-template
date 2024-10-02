# go-template
{{ cookiecutter.project_description }}

# how to run
local: `GIN_MODE=release go run main.go`
docker compose: `docker compose up -d`


## swagger setup
install go swagger cli: `go install github.com/swaggo/swag/cmd/swag@latest`
run `swag init` for document generation.
