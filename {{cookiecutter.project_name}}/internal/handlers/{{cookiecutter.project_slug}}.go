package handlers

import (
	"log"

	"github.com/atlasgames-sites/backend/internal"
	"github.com/gin-gonic/gin"
)

type {{cookiecutter.project_slug.title()}}Handler struct {
	DB *internal.DB
}

func (c *{{cookiecutter.project_slug.title()}}Handler) Get_{{cookiecutter.project_slug}}() gin.H {
	// get {{cookiecutter.project_slug}}
	log.Print(c.DB.Conn.Config().Host)
	return gin.H{"message": "done"}
}
