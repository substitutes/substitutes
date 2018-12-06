package routes

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/substitutes/substitutes/helpers"
)

// Root endpoint for listing all classes
func (ctl *Controller) List(c *gin.Context) {
	list, err := ctl.GetList()
	if err != nil {
		NewAPIError("Failed to fetch classes", err).Throw(c, 500)
		return
	}
	c.JSON(200, list)
}

func (ctl *Controller) GetList() ([]string, error) {
	resp, err := helpers.Request("Druck_Kla.htm")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
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
	return classes, nil
}
