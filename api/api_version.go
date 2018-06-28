package api

import (
	"github.com/gin-gonic/gin"
	"github.com/fronbasal/substitutes/helpers"
)

// Version endpoint for showing the current git commit and version history
func Version(c *gin.Context) {
	version, err := helpers.GetVersion()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
	c.JSON(200, version)
}
