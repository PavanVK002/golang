package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
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
	time.Sleep(5 * time.Second)

	// Extract data using JavaScript
	script := `
		var rows = document.querySelectorAll("table:has(thead tr td[colspan='12']) tbody tr");
		var data = [];
		for (var i = 0; i < rows.length; i++) {
			var rowData = rows[i].innerText;
			data.push(rowData);
		}
		return data;
	`
	var data []string
	if err := page.RunScript(script, nil, &data); err != nil {
		log.Fatalf("Failed to execute script: %v", err)
	}

	// Create a new Excel file
	f := excelize.NewFile()

	// Set sheet name
	sheetName := "BSE Data"
	index := f.NewSheet(sheetName)

	// Set header row
	headers := []string{"Index", "Open", "High", "Low", "Current Value", "Prev. Close", "Ch (pts)", "Ch (%)", "52 Wk High", "52 WK Low", "Turnover (Rs. Cr)", "% in Total Turnover"}
	for col, header := range headers {
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", excelize.ToAlphaString(col+1), 1), header)
	}

	// Split data and write to Excel file
	for rowIdx, row := range data {
		fields := strings.Fields(row)
		for colIdx, field := range fields {
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", excelize.ToAlphaString(colIdx+1), rowIdx+2), field)
		}
	}

	// Set active sheet
	f.SetActiveSheet(index)

	// Save Excel file
	if err := f.SaveAs("bse_data.xlsx"); err != nil {
		log.Fatalf("Failed to save Excel file: %v", err)
	}

	fmt.Println("Excel file created successfully.")
}
