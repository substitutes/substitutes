package main

import (
	"github.com/gin-gonic/gin"
	vapi "github.com/fronbasal/vertretungsplan/api"
)

// GinEngine returns an instance of the gin Engine.
func GinEngine() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("ui/*")

	r.Static("a", "a")

	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", nil) })

	r.GET("/k/:k", func(c *gin.Context) { c.HTML(200, "list.html", gin.H{"class": c.Param("k")}) })

	api := r.Group("api")
	{
		api.GET("/", vapi.Root)

		api.GET("/:class", vapi.Parser)
	}

	return r
}

func main() {
	GinEngine().Run(":5000")
}
