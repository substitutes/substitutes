package api

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"github.com/fronbasal/substitutes/structs"
	"strings"
)

// Version endpoint for showing the current git commit and version history
func Version(c *gin.Context) {
	b, err := exec.Command("git", "log", "-1", "--pretty=%B", "--oneline").CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	data := strings.Split(string(b[:]), " ")
	v := structs.Version{Hash: data[0], Message: strings.Replace(strings.Join(data[1:], " "), "\n", "", -1), Dirty: true}
	c.JSON(200, v)
}
