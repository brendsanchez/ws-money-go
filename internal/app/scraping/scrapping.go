package scraping

import (
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

func visitRoute(route string, c *colly.Collector) error {
	c.OnError(func(r *colly.Response, err error) {
		logrus.Error("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)
	})

	return c.Visit(route)
}
