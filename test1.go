package main

import (
    "fmt"
    "log"

    "github.com/gocolly/colly/v2"
)

func main() {
    url := "https://demo.opencart.com/"

    c := colly.NewCollector()
	c.OnHTML("div.col", func(e *colly.HTMLElement) {
        var image,productName,productLink,description,oldPrice,newPrice,tax string
        // image := e.ChildAttr("div.image a", "href")
        // productName := e.ChildText("h4 a")
        // productLink := e.ChildAttr("h4 a", "href")
        // description := e.ChildText("p")
        // oldPrice := e.ChildText(".price-old")
        // newPrice := e.ChildText(".price-new")
        // tax := e.ChildText(".price-tax")

        fmt.Println("Image:", image)
        fmt.Println("Product Name:", productName)
        fmt.Println("Product Link:", productLink)
        fmt.Println("Description:", description)
        fmt.Println("Old Price:", oldPrice)
        fmt.Println("New Price:", newPrice)
        fmt.Println("Tax:", tax)
        fmt.Println()
    })
	err := c.Visit(url)
    if err != nil {
        log.Fatal(err)
    }
}
