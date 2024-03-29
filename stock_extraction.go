package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"

    "github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
    url, image, name, price string
}

func main() {
    // initializing the slice of structs to store the data to scrape
    var pokemonProducts []PokemonProduct

    // creating a new Colly instance
    c := colly.NewCollector()

    // visiting the target page
    c.OnHTML("li.product", func(e *colly.HTMLElement) {
        pokemonProduct := PokemonProduct{}

        pokemonProduct.url = e.ChildAttr("a", "href")
        pokemonProduct.image = e.ChildAttr("img", "src")
        pokemonProduct.name = e.ChildText("h2")
        pokemonProduct.price = e.ChildText(".price")

        pokemonProducts = append(pokemonProducts, pokemonProduct)
    })

    // visiting the target page
    c.Visit("https://www.bseindia.com/Sensex/IndexHighlight.html")

    // Print the scraped data
    fmt.Println("Scraped data:")
    for _, pokemonProduct := range pokemonProducts {
        fmt.Printf("URL: %s, Image: %s, Name: %s, Price: %s\n", pokemonProduct.url, pokemonProduct.image, pokemonProduct.name, pokemonProduct.price)
    }

    // opening the CSV file
    file, err := os.Create("bse_stock_data.csv")
    if err != nil {
        log.Fatalln("Failed to create output CSV file", err)
    }
    defer file.Close()

    // initializing a file writer
    writer := csv.NewWriter(file)

    // writing the CSV headers
    headers := []string{
        "url",
        "image",
        "name",
        "price",
    }
    writer.Write(headers)

    // writing each Pokemon product as a CSV row
    for _, pokemonProduct := range pokemonProducts {
        // converting a PokemonProduct to an array of strings
        record := []string{
            pokemonProduct.url,
            pokemonProduct.image,
            pokemonProduct.name,
            pokemonProduct.price,
        }

        // adding a CSV record to the output file
        writer.Write(record)
    }
    defer writer.Flush()

    fmt.Println("Data written to CSV file successfully!")
}
