package routes

import (
	"log"

	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal/handlers"
	"github.com/gin-gonic/gin"
)

func {{cookiecutter.resource_name.title()}}Router(app *gin.Engine) {
	v1 := app.Group("/v1/{{cookiecutter.resource_name}}")
	{
		v1.GET("", get_{{cookiecutter.resource_name}})
	}
}

// get_{{cookiecutter.resource_name}}
// @Summery get {{cookiecutter.resource_name}}
// @Schemes
// @Description get {{cookiecutter.resource_name}}
// @Tags {{cookiecutter.resource_name}}_v1
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /v1/{{cookiecutter.resource_name}} [get]
func get_{{cookiecutter.resource_name}}(ctx *gin.Context) {
	// load {{cookiecutter.resource_name}} handler
	handler, _ := ctx.Get("{{cookiecutter.resource_name}}_handler")
	{{cookiecutter.resource_name}}_handler := handler.(handlers.{{cookiecutter.resource_name.title()}}Handler)

	log.Print({{cookiecutter.resource_name}}_handler.DB.Conn.Config())
	ctx.JSON(200, gin.H{"message": "done"})
}
