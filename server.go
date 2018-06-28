package main

import (
	vapi "github.com/fronbasal/substitutes/api"
	"github.com/gin-gonic/gin"
	"github.com/fronbasal/substitutes/helpers"
)

// GinEngine returns an instance of the gin Engine.
func GinEngine() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("ui/*")

	r.Static("a", "a")

	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", gin.H{"version": helpers.GetVersionString()}) })

	r.GET("/c/:c", func(c *gin.Context) { c.HTML(200, "list.html", gin.H{"class": c.Param("c"), "version": helpers.GetVersionString()}) })

	api := r.Group("api")
	{
		api.GET("/", vapi.Root)

		api.GET("/c/:class", vapi.Parser)

		api.GET("/t/:teacher", vapi.Teacher)

		api.GET("/version", vapi.Version)
	}

	return r
}

func main() {
	GinEngine().Run(":5000")
}
