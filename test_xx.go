package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)
var category = []string{"Market_Cap_Broad", "Sector&Industry", "Thematics", "Startegy", "Sustainability", "Volatility", "Composite", "Government", "Corporate", "Money_Market"}

func main() {
	start := time.Now()
	fmt.Printf("Start time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	for i:=10;i>0;i--{
		url := fmt.Sprintf("https://www.bseindia.com/Sensex/IndexHighlight.html?type=%d",i)
		filename :=fmt.Sprintf("BSE_data_%s_001.json",category[i-1])
		Webscrapping_IndexHighlight(url,filename)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
		// Print current time of execution
	fmt.Printf("End time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	time.Sleep(1* time.Minute)
}


func Webscrapping_IndexHighlight(url string ,filename string){

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
	time.Sleep(10 * time.Second)

	var data [][]string
	// fmt.Println(fetchData(page, "6"))
	// Extract data using JavaScript
	var columns = []string{"12", "6"}
	for _,colom := range columns{
		y :=   fetchData(page, colom)
		if len(y)>0{
			data = append(data,y)
		}
	}

	Data_Extract_Save(filename,data)

}
///////////////////////////////////////////////////////////////////////////////////////////////////////////
func fetchData(page *agouti.Page,col string,)[]string{
	script := fmt.Sprintf(`
		var rows = document.querySelectorAll("table:has(thead tr td[colspan='%s']) tbody tr");
		var data = [];
		var cells = rows[0].querySelectorAll("td");
		data.push(cells[0].innerText);
		return data;
	`, col)
	var data []string
	if err := page.RunScript(script, nil, &data); err != nil {
		return data
		// log.Fatalf("Failed to execute script: %v", err)
	}
	return data
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////
func Data_Extract_Save(filename  string ,data [][]string){
	// Create a new JSON file
	file, err := os.Create("json_folder_test/"+filename)
		if err != nil {
			log.Fatalf("Failed to create JSON file: %v", err)
		}
		defer file.Close()
	
		// Marshal data to JSON and write to the file
		encoder := json.NewEncoder(file)
		
		alltable  := make([][]map[string]string,0)
		for _, table := range data {
			// Extract table content
			content := table[0]
			// fmt.Println(table[0])
			// Split the content by newline to get individual rows
			rows := strings.Split(content, "\n")
			
			// Extract headers from the first row
			headers := strings.Split(rows[1],"\t")
			// for i,value := range  headers {
			// 	fmt.Printf("%d,  %s\n", i, value)
			// }
			// 	fmt.Println(rows[2])
			var tableData []map[string]string
			for _, row := range rows[2:] {
				fields := strings.Split(row,"\t")
				rowMap := make(map[string]string)
				for j, field := range fields {
					// Use headers as keys
					rowMap[headers[j]] = field
					// fmt.Println(rows[2])
				}
				// fmt.Println(rowMap)
				tableData = append(tableData, rowMap)
			}
			alltable = append(alltable, tableData)
			// Convert table data to JSON and write to the file
			
		}
		if err := encoder.Encode(alltable); err != nil {
			log.Printf("Error writing JSON data: %v", err)
		}
		fmt.Printf("JSON file for %s created successfully .\n", filename)
	
	}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
