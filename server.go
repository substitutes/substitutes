package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	vapi "github.com/substitutes/substitutes/api"
	"github.com/substitutes/substitutes/helpers"
)

// GinEngine returns an instance of the gin Engine.
func GinEngine() *gin.Engine {

	r := gin.Default()

	// Create the custom multitemplate renderer
	r.HTMLRender = newRenderer()

	// Create the static directory and path
	r.Static("a", "a")

	// Create the index view
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{"version": helpers.GetVersionString()})
	})

	r.GET("/c/:c", func(c *gin.Context) {
		c.HTML(200, "list", gin.H{"class": c.Param("c"), "version": helpers.GetVersionString()})
	})

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
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("substitutes")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		// Environment variables are also allowed
		log.Warn("Failed to read configuration file: ", err)
	}
	log.Info("Initialized application")

	GinEngine().Run(":5000")
}

func newRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "ui/base.html", "ui/index.html")
	r.AddFromFiles("list", "ui/base.html", "ui/list.html")
	return r
}
