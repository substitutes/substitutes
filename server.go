package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
	"github.com/gin-gonic/gin"
)

type vertretung struct {
	Class     string
	Std       string
	Teacher   string
	Subject   string
	Room      string
	Type      string
	Notes     string
	Cancelled bool
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("ui/*")

	r.Static("a", "a")

	r.GET("/", func(c *gin.Context) { c.HTML(200, "index.html", nil) })

	r.GET("/k/:k", func(c *gin.Context) { c.HTML(200, "list.html", gin.H{"klasse": c.Param("k")}) })

	api := r.Group("api")
	{
		api.GET("/", func(c *gin.Context) {
			resp, err := request("Druck_Kla.htm")
			if err != nil {
				c.JSON(500, gin.H{"message": "Failed to make request", "error": err.Error()})
				return
			}
			defer resp.Body.Close()
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				c.JSON(500, gin.H{"message": "Failed to read document", "error": err.Error()})
				return
			}
			var classes []string
			doc.Find("table").Last().Find("td").Each(func(i int, sel *goquery.Selection) {
				title := sel.Text()
				if title != "" {
					if title == "---" {
						title = "entfall"
					}
					classes = append(classes, title)
				}
			})
			c.JSON(200, classes)
		})

		api.GET("/:klasse", func(c *gin.Context) {
			k := c.Param("klasse")
			if k == "entfall" {
				k = "___"
			} else if !regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString(k) {
				c.JSON(400, gin.H{"message": "Invalid class!"})
				return
			}
			resp, err := request("Druck_Kla_" + k + ".htm")
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
			var vertretungen []vertretung
			doc.Find("table").Last().Remove()
			doc.Find("table").Last().Find("tr").Each(func(i int, sel *goquery.Selection) {
				if i != 0 {
					var v vertretung
					sel.Find("td font").Each(func(i int, sel *goquery.Selection) {
						t := strings.Replace(sel.Text(), "\n", "", -1)
						switch i {
						// dis ugly bc fuck html
						case 0:
							v.Class = sel.Find("b").Text()
							break
						case 1:
							v.Std = t
							break
						case 2:
							v.Teacher = strings.Replace(t, "?", " => ", 1)
							break
						case 3:
							if strings.Contains(t, "R") {
								v.Subject = ""
							} else {
								v.Subject = t
							}
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
					vertretungen = append(vertretungen, v)
				}
			})

			var meta struct {
				Datum  string
				Klasse string
			}
			meta.Datum = strings.Replace(doc.Find("center font font b").First().Text(), "\n", "", -1)
			meta.Klasse = strings.Replace(doc.Find("center font font font").First().Text(), "\n", "", -1)
			c.JSON(200, gin.H{"vertretungen": vertretungen, "meta": meta})
		})
	}

	r.Run(":5000")
}

func request(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://www.stgym.de/ovp/"+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c().Username, c().Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func c() Credentials {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatal("Failed to read config file!")
	}
	var c Credentials
	json.Unmarshal(b, &c)
	return c
}
