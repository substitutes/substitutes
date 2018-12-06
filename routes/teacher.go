package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/structs"
	"strings"
)

// Teacher endpoint for the teacher view
func (ctl *Controller) Teacher(c *gin.Context) {
	classes, err := ctl.GetList()
	if err != nil {
		NewAPIError("Failed to fetch classes", err).Throw(c, 500)
		return
	}

	// Get the whole dataset
	var responses []structs.SubstituteResponse
	for i := range classes {
		response, err := ctl.GetClass(classes[i])
		if err != nil {
			err.Throw(c, 500)
			return
		}
		responses = append(responses, response)
	}

	teacher := c.Param("teacher")
	var matches []structs.Substitute

	for i := range responses {
		for n := range responses[i].Substitutes {
			if strings.ToLower(responses[i].Substitutes[n].Teacher) == strings.ToLower(teacher) {
				matches = append(matches, responses[i].Substitutes[n])
			}
		}
	}

	c.JSON(200, matches)
}
