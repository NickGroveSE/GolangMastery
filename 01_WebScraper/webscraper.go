package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	// filmNames := []string{}

	fmt.Print("Instantiating...\n")
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("letterboxd.com"),
	)

	c.OnHTML(".image", func(e *colly.HTMLElement) {
		// film := e.Attr("data-film-name")
		// filmNames = append(filmNames, film)
		fmt.Println(e.Attr("alt"))
	})

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println(filmNames)
	// })

	c.Visit("https://letterboxd.com/feauv/")

}
