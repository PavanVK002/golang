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
	
	
	Data_Extract_Save(Webscrapping_IndexHighlight("https://www.bseindia.com/Sensex/IndexHighlight.html"),"bse_data1.csv")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func Webscrapping_IndexHighlight(url string) []string{

	// Start a WebDriver session
	driver := agouti.ChromeDriver(
		// You may need to specify the path to your ChromeDriver executable
		agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox","--log-level=3"}),
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
	if err := page.Navigate(url); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Wait for dynamic content to load (adjust time as needed)
	time.Sleep(2 * time.Second)

	// Extract data using JavaScript
	script := `
	var rows = document.querySelectorAll("table:has(thead tr td[colspan='6']) tbody tr");
	var data = [];
	var cells = rows[0].querySelectorAll("td");
	data.push(cells[0].innerText);
	return data;
	`
	var data []string
	if err := page.RunScript(script, nil, &data); err != nil {
		log.Fatalf("Failed to execute script: %v", err)
	}
	fmt.Println(data)
	return data
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////
func Data_Extract_Save(data []string,filename  string){
	rows := strings.Split(data[0], "\n")
	
	// Create a new CSV file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	// // Write the data to the CSV file
	writer := csv.NewWriter(file)
	defer writer.Flush()


	for _, row := range rows {
		
			fields := strings.Split(row, "\t")
			if err := writeRowToFile(file, fields); err != nil {
				log.Printf("Error writing row to CSV: %v", err)
			}
	
	}
	fmt.Println("CSV file created successfully.")
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func writeRowToFile(file *os.File, fields []string) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()
	return writer.Write(fields)
}