package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/helpers"
)

// Version endpoint for showing the current git commit and version history
func (ctl *Controller) Version(c *gin.Context) {
	version, err := helpers.GetVersion()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
	c.JSON(200, version)
}
