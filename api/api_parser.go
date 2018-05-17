package api

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"github.com/fronbasal/substitutes/helpers"
	"github.com/fronbasal/substitutes/structs"
	"github.com/gin-gonic/gin"
)

// Parser function for returning the endpoint at /api/c/{class}
func Parser(c *gin.Context) {
	k := c.Param("class")
	if k == "Cancelled" {
		k = "___"
	} else if !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(k) {
		c.JSON(400, gin.H{"message": "Invalid class!"})
		return
	}
	resp, err := helpers.Request("Druck_Kla_" + k + ".htm")
	if resp.StatusCode == 404 {
		c.JSON(404, gin.H{"message": "Not found."})
		return
	}
	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{"message": "Expected 200, got: " + resp.Status})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to make request", "error": err.Error()})
		return
	}
	defer resp.Body.Close()

	utfBody, err := iconv.NewReader(resp.Body, "iso-8859-1", "utf-8")
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to decompose UTF8"})
		return
	}

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to read document", "error": err.Error()})
		return
	}
	var substitutes []structs.Substitutes
	doc.Find("table").Last().Remove()
	doc.Find("table").Last().Find("tr").Each(func(i int, sel *goquery.Selection) {
		if i != 0 {
			var v structs.Substitutes
			count := len(sel.Find("td").Nodes)
			if count == 15 {
				sel.Find("td").Each(func(i int, sel *goquery.Selection) {
					t := strings.Replace(sel.Text(), "\n", "", -1)
					t = strings.TrimSpace(t)
					switch i {
					// Parse the HTML table into the struct
					case 0:
						v.Date = t
						break
					case 1:
						v.Hour = t
						break
					case 2:
						v.Day = t
						break
					case 3:
						v.Teacher = t
						break
					case 4:
						v.Time = t
						break
					case 5:
						v.Subject = t
						break
					case 6:
						v.Type = t
						break
					case 7:
						v.Notes = t
						break
					case 8:
						v.Classes = t
						break
					case 9:
						v.Room = strings.Replace(t, "?", " => ", 1)
						break
					case 10:
						v.After = t
						break
					case 11:
						// Check if there is content
						v.Cancelled = len(strings.Replace(t, " ", "", -1)) != 0
						break
					case 12:
						matched, err := regexp.MatchString("x|X", t)
						if err != nil {
							c.JSON(500, gin.H{"message": "Failed to compile Regex."})
							return
						}
						v.New = matched
						break
					case 13:
						v.Reason = t
						break
					case 14:
						v.Counter = t
						break
					}
				})
			} else { // Alternative parser, deprecated
				sel.Find("td font").Each(func(i int, sel *goquery.Selection) {
					t := strings.Replace(sel.Text(), "\n", "", -1)
					switch i {
					case 0:
						v.Classes = sel.Find("b").Text()
						break
					case 1:
						v.Hour = t
						break
					case 2:
						v.Teacher = strings.Replace(t, "?", " => ", 1)
						break
					case 3:
						v.Subject = t
						break
					case 4:
						v.Room = strings.Replace(t, "?", " => ", 1)
						break
					case 5:
						v.Type = t
						break
					case 6:
						v.Notes += t
						break
					}
				})

			}
			substitutes = append(substitutes, v)

		}
	})

	var meta struct {
		Date  string `json:"date"`
		Class string `json:"class"`
	}
	meta.Date = strings.Replace(doc.Find("center font font b").First().Text(), "\n", "", -1)
	meta.Class = strings.Replace(doc.Find("center font font font").First().Text(), "\n", "", -1)
	c.JSON(200, gin.H{"substitutes": substitutes, "meta": meta})
}
