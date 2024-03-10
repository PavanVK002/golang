package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	// Start a WebDriver session
	driver := agouti.ChromeDriver(
		// You may need to specify the path to your ChromeDriver executable
		agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox"}),
	)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	// Open a new page
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page: %v", err)
	}

	// Navigate to the BSE India website
	if err := page.Navigate("https://www.bseindia.com/Sensex/IndexHighlight.html"); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Wait for dynamic content to load (adjust time as needed)
	time.Sleep(2 * time.Second)

	// Extract data using JavaScript
	script := `
	var rows = document.querySelectorAll("table:has(thead tr td[colspan='12']) tbody tr");
	var data = [];
	var cells = rows[0].querySelectorAll("td");
	data.push(cells[0].innerText);
	return data;
	`
	var data []string
	if err := page.RunScript(script, nil, &data); err != nil {
		log.Fatalf("Failed to execute script: %v", err)
	}
	//fmt.Println(data)
	rows := strings.Split(data[0], "\n")
	//fmt.Println(data[0][0])
	
	// for i, row := range rows {
	// 	fmt.Printf("Row %d: %v\n", i+1, row)
	// }
	// Create a new CSV file
	file, err := os.Create("bse_data1.csv")
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	// // Write the data to the CSV file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range rows {
		fields := strings.Split(row, "\t")
		if err := writer.Write(fields); err != nil {
			log.Fatalf("Error writing row to CSV: %v", err)
		}
	}

	fmt.Println("CSV file created successfully.")
}
