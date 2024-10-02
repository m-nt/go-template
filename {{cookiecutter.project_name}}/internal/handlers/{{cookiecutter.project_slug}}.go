package handlers

import (
	"log"

	"github.com/{{cookiecutter.project_github}}/{{cookiecutter.project_name}}/internal"
	"github.com/gin-gonic/gin"
)

type {{cookiecutter.project_name.title()}}Handler struct {
	DB *internal.DB
}

func (c *{{cookiecutter.project_name.title()}}Handler) Get_{{cookiecutter.project_slug}}() gin.H {
	// get {{cookiecutter.project_name}}
	log.Print(c.DB.Conn.Config().Host)
	return gin.H{"message": "done"}
}
