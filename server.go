package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/substitutes/substitutes/helpers"
	"github.com/substitutes/substitutes/lookup"
	vapi "github.com/substitutes/substitutes/routes"
	"strconv"
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

	r.GET("/t/", func(c *gin.Context) {
		c.HTML(200, "teacherlist", gin.H{"version": helpers.GetVersionString()})
	})

	r.GET("/t/:t", func(c *gin.Context) {
		c.HTML(200, "teacherview", gin.H{"teacher": c.Param("t"), "version": helpers.GetVersionString()})
	})

	ctl := vapi.NewController()

	api := r.Group("api")
	{
		api.GET("/", ctl.List)

		api.GET("/c/:class", ctl.Parser)

		api.GET("/t/", ctl.ListTeachers)

		api.GET("/t/:teacher", ctl.Teacher)

		api.GET("/version", ctl.Version)
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

	lookup.New().ReadFile()

	viper.SetDefault("port", 5000)
	GinEngine().Run(":" + strconv.Itoa(viper.GetInt("port")))
}

func newRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "ui/base.html", "ui/index.html")
	r.AddFromFiles("list", "ui/base.html", "ui/list.html")

	r.AddFromFiles("teacherlist", "ui/base.html", "ui/teacherlist.html")
	r.AddFromFiles("teacherview", "ui/base.html", "ui/teacherview.html")
	return r
}
