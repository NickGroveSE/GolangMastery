package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	favorites := []string{}
	recents := []string{}
	ratings := []string{}

	fmt.Print("Instantiating...\n\n")

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("letterboxd.com"),
	)

	c.OnHTML("#favourites", func(e *colly.HTMLElement) {
		favorites = e.ChildAttrs("ul > li > div.poster > img", "alt")

	})

	c.OnHTML("#recent-activity", func(e *colly.HTMLElement) {
		recents = e.ChildAttrs("ul > li > div.poster > img", "alt")
		ratings = strings.Split(e.ChildText("ul > li > p > span"), " ")
	})

	c.OnScraped(func(r *colly.Response) {

		printFaves(favorites)

		printRecents(recents, ratings)

		fmt.Println(ratings[1])

	})

	c.Visit("https://letterboxd.com/filmarchive/")

}

func printFaves(films []string) {

	fmt.Print("Favorite Films:\n")

	for _, film := range films {
		fmt.Println("\t" + film)
	}

	fmt.Print("\n")
}

func printRecents(films []string, ratings []string) {

	fmt.Print("Recent Films:\n")

	for index, film := range films {
		fmt.Println("\t" + film + " - " + ratings[index+index])
	}

	fmt.Print("\n")

}
