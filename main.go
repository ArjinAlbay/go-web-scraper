package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {


	
	var giveMeLink string
	fmt.Print("Enter the link (without https://): ")
	fmt.Scanln(&giveMeLink)

	// Add "https://" before scanning the input
	
	fmt.Println("Full link:", giveMeLink)



	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only the specified domain
		colly.AllowedDomains(giveMeLink ),
	)
	
	
	
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        href := e.Attr("href")
        fmt.Println(href)
    })

	// Log each request
	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting: %s", r.URL.String())
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %v failed with response: %v. Error: %v", r.Request.URL, r, err)
	})


	// Add "https://" before the input link
	giveMeLink = "https://" + giveMeLink + "/"

	
	// Start scraping
	if err := c.Visit(giveMeLink + "/"); err != nil {
		log.Fatalf("Failed to visit the website: %v", err)
	}
}