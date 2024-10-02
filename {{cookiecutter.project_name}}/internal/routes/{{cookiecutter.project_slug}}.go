package routes

import (
	"log"

	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal/handlers"
	"github.com/gin-gonic/gin"
)

func {{cookiecutter.project_name.title()}}Router(app *gin.Engine) {
	v1 := app.Group("/v1/{{cookiecutter.project_name}}")
	{
		v1.GET("", get_{{cookiecutter.project_name}})
	}
}

// get_{{cookiecutter.project_name}}
// @Summery get {{cookiecutter.project_name}}
// @Schemes
// @Description get {{cookiecutter.project_name}}
// @Tags {{cookiecutter.project_name}}_v1
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /v1/{{cookiecutter.project_name}} [get]
func get_{{cookiecutter.project_name}}(ctx *gin.Context) {
	// load {{cookiecutter.project_name}} handler
	handler, _ := ctx.Get("{{cookiecutter.project_name}}_handler")
	{{cookiecutter.project_name}}_handler := handler.(handlers.{{cookiecutter.project_name.title()}}Handler)

	log.Print({{cookiecutter.project_name}}_handler.DB.Conn.Config())
	ctx.JSON(200, gin.H{"message": "done"})
}
