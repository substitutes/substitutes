package main

import (
	vapi "github.com/substitutes/substitutes/api"
	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/helpers"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"time"
)

var (
	verbose = kingpin.Flag("verbose", "Enable verbose output").Short('v').Bool()
)

// GinEngine returns an instance of the gin Engine.
func GinEngine() *gin.Engine {
	kingpin.Parse()

	logrus.SetLevel(logrus.WarnLevel)
	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	r := gin.New()

	r.Use(gin.Recovery(), ginrus.Ginrus(logrus.StandardLogger(), time.RFC822, true))

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

	logrus.Debug("Initialized application.")

	return r
}

func main() {
	GinEngine().Run(":5000")
}
