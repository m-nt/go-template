package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/docs"
	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal"
	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal/middlewares"
	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal/routes"
)

// @BasePath /

// liveness
// @Summery liveness
// @Schemes
// @Description liveness
// @Tags liveness
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /liveness [get]
func liveness(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "ok"})
}

func main() {
	_settings := internal.Get_Settings()
	log.Print(_settings)
	// app initialization
	app := gin.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"x-user-*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// middlewares
	app.Use(internal.DefaultStructuredLogger())
	app.Use(gin.Recovery())

	// user middlewares
	app.Use(middlewares.Authenticate())

	// database setup
	db := internal.Get_database()
	defer internal.Close_database(&db)
	app.Use(middlewares.AddHandlers(&db))

	// user routes
	app.GET("/liveness", liveness)

	routes.UserRouter(app)

	// swagger setup
	docs.SwaggerInfo.BasePath = "/"
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// running server
	app.Run(_settings.SERVICE_ADDRESS)
}
