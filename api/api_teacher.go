package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fronbasal/substitutes/structs"
	"github.com/gin-gonic/gin"
)

// Teacher endpoint for the teacher view
func Teacher(c *gin.Context) {
	// TODO: Variable host
	resp, err := http.Get("http://localhost:5000/api")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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
		Substitutes []structs.Substitutes `json:"substitutes"`
	}

	type multiResponse struct {
		Meta struct {
			Date  string `json:"date"`
			Class string `json:"class"`
		} `json:"meta"`
		Substitute structs.Substitutes `json:"substitute"`
	}

	var teachers []multiResponse

	for _, class := range v {
		apiResp, err := http.Get("http://localhost:5000/api/c/" + class)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var r response
		if err := json.NewDecoder(apiResp.Body).Decode(&r); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
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
