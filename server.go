package main

import (
	vapi "github.com/fronbasal/substitutes/api"
	"github.com/gin-gonic/gin"
)

// GinEngine returns an instance of the gin Engine.
func GinEngine() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("ui/*")

	r.Static("a", "a")

	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", nil) })

	r.GET("/k/:k", func(c *gin.Context) { c.HTML(200, "list.html", gin.H{"class": c.Param("k")}) })

	r.GET("/t/:teacher", func(c *gin.Context) {
		c.HTML(200, "teacher.html", gin.H{"teacher": c.Param("teacher")})
	})

	api := r.Group("api")
	{
		api.GET("/", vapi.Root)

		api.GET("/c/:class", vapi.Parser)

		api.GET("/t/:teacher", vapi.Teacher)
	}

	return r
}

func main() {
	GinEngine().Run(":5000")
}
