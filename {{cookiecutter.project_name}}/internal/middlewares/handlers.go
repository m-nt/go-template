package middlewares

import (
	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal"
	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AddHandlers(db_conn *internal.DB) gin.HandlerFunc {
	user_handler := handlers.{{cookiecutter.project_slug.title()}}Handler{DB: db_conn}
	// add other handlers here and set them in context

	return func(ctx *gin.Context) {
		ctx.Set("{{cookiecutter.project_slug}}_handler", {{cookiecutter.project_slug}}_handler)
	}
}
