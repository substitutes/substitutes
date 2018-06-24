package api

import (
	"github.com/gin-gonic/gin"
	"github.com/fronbasal/substitutes/helpers"
)

func Auth(c *gin.Context) {
	auth, err := helpers.IServLogin(c.PostForm("username"), c.PostForm("password"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"auth": auth})
}
