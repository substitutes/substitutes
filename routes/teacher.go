package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/substitutes/substitutes/structs"
	"reflect"
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
			t := responses[i].Substitutes[n].Teacher
			if strings.Contains(t, "=>") {
				// Sanitize, only get newer
				t = strings.Split(t, " => ")[1]
			}

			if strings.ToLower(t) == strings.ToLower(teacher) {
				matched := false
				for x := range matches {
					if reflect.DeepEqual(matches[x], responses[i].Substitutes[n]) {
						matched = true
					}
				}
				if !matched {
					matches = append(matches, responses[i].Substitutes[n])
				}
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
				teacher = strings.Split(teacher, " => ")[1]
			}
			if len(teacher) < 2 {
				continue
			}
			if !stringSliceContains(teachers, teacher) {
				teachers = append(teachers, teacher)
			}
		}
	}
	if len(teachers) == 0 {
		NewAPIError("No substitutes for "+c.Param("teacher"), errors.New("no content")).Throw(c, 204)
		return
	}

	c.JSON(200, teachers)
}

func stringSliceContains(s []string, e string) bool {
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
