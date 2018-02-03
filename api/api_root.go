package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/fronbasal/substitutes/helpers"
)

func Root(c *gin.Context) {
	resp, err := helpers.Request("Druck_Kla.htm")
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
				title = "Cancelled"
			} else if title == "XXX" {
				title = "Break Supervisor"
			}
			classes = append(classes, title)
		}
	})
	c.JSON(200, classes)
}
