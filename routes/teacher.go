package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/structs"
	"strings"
)

// Teacher endpoint for the teacher view
func (ctl *Controller) Teacher(c *gin.Context) {
	responses, err := ctl.GetAll()

	if err != nil {
		err.Throw(c, 500)
		return
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

func (ctl *Controller) ListTeachers(c *gin.Context) {
	responses, err := ctl.GetAll()
	if err != nil {
		err.Throw(c, 500)
		return
	}

	var teachers []string
	for i := range responses {
		for n := range responses[i].Substitutes {
			teacher := responses[i].Substitutes[n].Teacher
			if strings.Contains(teacher, "=>") {
				// Sanitize, only get newer
				teacher = strings.Split(responses[i].Substitutes[n].Teacher, " => ")[1]
			}
			if len(teacher) < 2 {
				continue
			}
			if !contains(teachers, teacher) {
				teachers = append(teachers, teacher)
			}
		}
	}

	c.JSON(200, teachers)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (ctl *Controller) GetAll() ([]structs.SubstituteResponse, *APIErrorMessage) {
	classes, err := ctl.GetList()
	if err != nil {
		return nil, NewAPIError("Failed to fetch classes", err)
	}

	// Get the whole dataset
	var responses []structs.SubstituteResponse
	for i := range classes {
		response, err := ctl.GetClass(classes[i])
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}
