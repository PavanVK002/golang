package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/sclevine/agouti"
)

var category = []string{"Market_Cap_Broad", "Sector&Industry", "Thematics", "Startegy", "Sustainability", "Volatility", "Composite", "Government", "Corporate", "Money_Market"}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Webscrapping_IndexHighlight(i)
		}(i)
	}

	wg.Wait()
}

func Webscrapping_IndexHighlight(type_ int) {
	url := fmt.Sprintf("https://www.bseindia.com/Sensex/IndexHighlight.html?type=%d", type_)
	filename := fmt.Sprintf("BSE_Data_%s.json", category[type_-1])

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu", "--no-sandbox", "--log-level=3"}),
	)
	if err := driver.Start(); err != nil {
		log.Printf("Failed to start driver: %v", err)
		return
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Printf("Failed to open page: %v", err)
		return
	}

	if err := page.Navigate(url); err != nil {
		log.Printf("Failed to navigate: %v", err)
		return
	}

	time.Sleep(2 * time.Second)

	var data [][]string
	columns := []string{"12", "6"}
	count := 2
	for _, colom := range columns {
		y := fetchData(page, colom)
		if len(y) > 0 {
			data = append(data, y)
		} else {
			count--
		}
	}
	if count == 0 {
		log.Printf("Data for %s did not fetch \n", filename)
		return
	}
	Data_Extract_Save(filename, data)
}

func fetchData(page *agouti.Page, col string) []string {
	script := fmt.Sprintf(`
		var rows = document.querySelectorAll("table:has(thead tr td[colspan='%s']) tbody tr");
		var data = [];
		var cells = rows[0].querySelectorAll("td");
		data.push(cells[0].innerText);
		return data;
	`, col)
	var data []string
	if err := page.RunScript(script, nil, &data); err != nil {
		log.Printf("Failed to execute script: %v", err)
	}
	return data
}

func Data_Extract_Save(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Failed to create JSON file: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	alltable := make([][]map[string]string, 0)
	for _, table := range data {
		content := table[0]
		rows := strings.Split(content, "\n")

		headers := strings.Split(rows[1], "\t")

		var tableData []map[string]string
		for _, row := range rows[2:] {
			fields := strings.Split(row, "\t")
			rowMap := make(map[string]string)
			for j, field := range fields {
				rowMap[headers[j]] = field
			}
			tableData = append(tableData, rowMap)
		}
		alltable = append(alltable, tableData)
	}

	if err := encoder.Encode(alltable); err != nil {
		log.Printf("Error writing JSON data: %v", err)
	}
	log.Printf("JSON file for %s created successfully.\n", filename)
}
