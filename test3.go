package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

func main() {
	// Create a WebDriver session
	driver := agouti.ChromeDriver()
	err := driver.Start()
	if err != nil {
		log.Fatalf("Failed to start ChromeDriver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page: %v", err)
	}

	// Navigate to the website URL
	err = page.Navigate("https://www.bseindia.com/Sensex/IndexHighlight.html")
	if err != nil {
		log.Fatalf("Failed to navigate to website: %v", err)
	}

	// Wait for the dynamic content to load (adjust timeout as needed)
	err = page.Session().SetImplicitWait(50)
	if err != nil {
		log.Fatalf("Failed to set implicit wait: %v", err)
	}

	// Extract the HTML content after the page has fully loaded
	html, err := page.HTML()
	if err != nil {
		log.Fatalf("Failed to get HTML content: %v", err)
	}


	// Load the HTML content into goquery document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal("Error loading HTML content:", err)
	}

	// Find tables with the specified header
	tables := doc.Find("table:has(thead tr td[colspan='12'])")
	//fmt.Println(tables.Html())
		rows := tables.Find("tbody tr td tbody tr")

		// Iterate over each row in the table body
		rows.Each(func(j int, row *goquery.Selection) {
			// Extract the text content of each cell in the row
			cells := row.Find("td")
			cells.Each(func(k int, cell *goquery.Selection) {
				fmt.Printf("Table 1, Row %d, Column %d: %s\n",  j+1, k+1, cell.Text())
			})
			fmt.Println("-------------------------------------")
		})


}