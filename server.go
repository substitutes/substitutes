package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type class struct {
	Link  string
	Title string
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		resp, err := request("Druck_Kla.htm")
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to make request", "error": err.Error()})
		}
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to read document", "error": err.Error()})
		}
		var classes []class
		doc.Find("table").Last().Find("td").Each(func(i int, sel *goquery.Selection) {
			href, _ := sel.Find("a").First().Attr("href")
			title := sel.Text()
			if title != "" {
				if title == "---" {
					title = "Entfall"
				}
				classes = append(classes, class{Link: href, Title: title})
			}
		})
		c.JSON(200, classes)
	})

	r.Run(":5000")
}

func request(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://www.stgym.de/ovp/"+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("schueler", "31ovp#18")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
