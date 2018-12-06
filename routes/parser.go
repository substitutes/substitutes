package routes

import (
	"errors"
	"github.com/substitutes/substitutes/lookup"
	"log"
	"regexp"
	"strings"

	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/helpers"
	"github.com/substitutes/substitutes/structs"
	"io/ioutil"
)

// Parser function for returning the endpoint at /routes/c/{class}
func (ctl *Controller) Parser(c *gin.Context) {
	k := c.Param("class")
	if k == "Cancelled" {
		k = "___"
	} else if !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(k) {
		NewAPIError("Invalid class", errors.New("class not valid")).Throw(c, 400)
		return
	}
	substitutes, errorMessage := ctl.GetClass(k)
	if errorMessage != nil {
		errorMessage.Throw(c, 500)
	}
	c.JSON(200, substitutes)
}

func (ctl *Controller) GetClass(class string) (structs.SubstituteResponse, *APIErrorMessage) {
	resp, err := helpers.Request("Druck_Kla_" + class + ".htm")

	if err != nil {
		return structs.SubstituteResponse{}, NewAPIError("Failed to make request", err)
	}

	// Defer after checking.
	defer resp.Body.Close()

	f, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return structs.SubstituteResponse{}, NewAPIError("Failed to read body", err)
	}
	if resp.StatusCode == 404 {
		return structs.SubstituteResponse{}, NewAPIError("Could not find site", nil)
	}
	if resp.StatusCode != 200 {
		return structs.SubstituteResponse{}, NewAPIError("Did not receive status 200", errors.New(resp.Status))
	}

	body := make([]byte, len(f))
	// TODO: Handle errors.
	iconv.Convert(f, body, "iso-8859-1", "utf-8")

	if err != nil {
		return structs.SubstituteResponse{}, NewAPIError("Failed to decompose UTF8", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return structs.SubstituteResponse{}, NewAPIError("Failed to read document", err)
	}
	var extended bool
	var substitutes []structs.Substitute
	doc.Find("table").Last().Remove()
	doc.Find("table").Last().Find("tr").Each(func(i int, sel *goquery.Selection) {
		if i != 0 {
			var v structs.Substitute
			count := len(sel.Find("td").Nodes)
			if count >= 10 /* Not working ,_, */ {
				extended = true
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
						v.Teacher = lookup.TeacherLookup.Get(t)
						v.TeacherInitials = t
						break
					case 4:
						v.Time = t
						break
					case 5:
						v.Subject = t
						break
					case 6:
						v.Type = t
						if t == "Vertretung" {
							v.Type = "Substitute"
						}
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
							log.Fatal("Failed to compile parser RegEx")
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
				extended = false
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
						v.Teacher = lookup.TeacherLookup.Get(strings.Replace(t, "?", " => ", 1))
						v.TeacherInitials = strings.Replace(t, "?", " => ", 1)
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

	meta := structs.SubstituteMeta{
		Extended: extended,
		Date:     strings.Replace(strings.Replace(doc.Find("center font font b").First().Text(), "\n", "", -1), "Vertretungen ", "Substitutes", 1),
		Class:    strings.Replace(doc.Find("center font font font").First().Text(), "\n", "", -1),
		Updated:  doc.Find("table").First().Find("tr").Last().Find("td").Last().Text(),
	}
	return structs.SubstituteResponse{Meta: meta, Substitutes: substitutes}, nil
}
