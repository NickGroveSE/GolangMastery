package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	// filmNames := []string{}

	fmt.Print("Instantiating...\n\n")

	// fmt.Print("Favorites:\n")
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("letterboxd.com"),
	)

	c.OnHTML("#favourites", func(e *colly.HTMLElement) {
		// film := e.Attr("data-film-name")
		// filmNames = append(filmNames, film)

		for _, film := range e.ChildAttrs("ul > li > div.poster > img", "alt") {
			fmt.Println("\t" + film)
		}

	})

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println(filmNames)
	// })

	c.Visit("https://letterboxd.com/feauv/")

}
