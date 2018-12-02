package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/structs"
)

// Teacher endpoint for the teacher view
func (ctl *Controller) Teacher(c *gin.Context) {
	resp, err := http.Get("http://localhost:5000/routes")
	if err != nil {
		NewAPIError("Failed to request API", err).Throw(c, 500)
		return
	}
	defer resp.Body.Close()
	var v []string
	json.NewDecoder(resp.Body).Decode(&v)
	type response struct {
		Meta struct {
			Date  string `json:"date"`
			Class string `json:"class"`
		} `json:"meta"`
		Substitutes []structs.Substitute `json:"substitutes"`
	}

	type multiResponse struct {
		Meta struct {
			Date  string `json:"date"`
			Class string `json:"class"`
		} `json:"meta"`
		Substitute structs.Substitute `json:"substitute"`
	}

	var teachers []multiResponse

	for _, class := range v {
		apiResp, err := http.Get("http://localhost:5000/routes/c/" + class)
		if err != nil {
			NewAPIError("Failed to request API", err).Throw(c, 500)
			return
		}
		var r response
		if err := json.NewDecoder(apiResp.Body).Decode(&r); err != nil {
			NewAPIError("Failed to decode API", err).Throw(c, 500)
			return
		}

		for _, substitute := range r.Substitutes {
			// TODO: Verify Teacher

			if substitute.Type == "Vertretung" && strings.Contains(substitute.Teacher, "=>") && strings.Contains(substitute.Teacher, c.Param("teacher")) {
				substitute.Teacher = strings.Split(substitute.Teacher, "=> ")[1]
				teachers = append(teachers, multiResponse{Meta: r.Meta, Substitute: substitute})
			}
		}
	}

	c.JSON(200, teachers)
}
