package routes

import (
	"log"

	"github.com/atlasgames-sites/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func {{cookiecutter.project_slug.title()}}Router(app *gin.Engine) {
	v1 := app.Group("/v1/{{cookiecutter.project_slug}}")
	{
		v1.GET("", get_{{cookiecutter.project_slug}})
	}
}

// get_{{cookiecutter.project_slug}}
// @Summery get {{cookiecutter.project_slug}}
// @Schemes
// @Description get {{cookiecutter.project_slug}}
// @Tags {{cookiecutter.project_slug}}_v1
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /v1/{{cookiecutter.project_slug}} [get]
func get_{{cookiecutter.project_slug}}(ctx *gin.Context) {
	// load {{cookiecutter.project_slug}} handler
	handler, _ := ctx.Get("{{cookiecutter.project_slug}}_handler")
	{{cookiecutter.project_slug}}_handler := handler.(handlers.{{cookiecutter.project_slug.title()}}Handler)

	log.Print({{cookiecutter.project_slug}}_handler.DB.Conn.Config())
	ctx.JSON(200, gin.H{"message": "done"})
}
