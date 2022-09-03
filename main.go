package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)


func main()  {
	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com"),
	)
	c.SetRequestTimeout(120 * time.Second)

	c.OnHTML("div.top-matter p.title", func (h *colly.HTMLElement)  {
		fmt.Println(h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})
	c.Visit("https://old.reddit.com/r/dankmemes/")
}