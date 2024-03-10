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

var category = []string{"Market_Cap_Broad","Sector&Industry","Thematics","Startegy","Sustainability","Volatility","Composite","Government","Corporate","Mone Market"}
func main() {
	dataCh := make(chan int)
	fmt.Printf("Start time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	for i := 0; i < 10000; i++ {
		go Webscrapping_IndexHighlight(dataCh)
	 }
	
	for i := 1; i <= 2; i++ {
		dataCh <- i
	}

	// Wait for goroutines to finish
	time.Sleep(2 * time.Minute)
	close(dataCh)
	fmt.Printf("End time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}


func Webscrapping_IndexHighlight(dataCh chan int){
	
	for type_ := range dataCh {
	url := fmt.Sprintf("https://www.bseindia.com/Sensex/IndexHighlight.html?type=%d",type_)
	fmt.Println(url)
	filename := fmt.Sprintf("BSE_Data_%s.json",category[int(type_)-1])
	// Start a WebDriver session
	driver := agouti.ChromeDriver(
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
	time.Sleep(30 * time.Second)

	var data [][]string
	var columns = []string{"12", "6"}
	count := 2
	for _,colom := range columns{
		y :=   fetchData(page, colom)
		if len(y)>0{
			data = append(data,y)
		}else{

			count--
		}
		
	}
	if count==0{
		fmt.Printf("Data for %s didnt not fetched \n",filename)
		return
	}
	Data_Extract_Save(filename,data)

}
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
		log.Fatalf("Failed to execute script: %v \n", err)
		
		return data
		//log.Fatalf("Failed to execute script: %v", err)

	}
	return data
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Data_Extract_Save(filename  string ,data [][]string){

	file, err := os.Create("json_folder/"+filename)
	if err != nil {
		log.Fatalf("Failed to create JSON file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	
	alltable  := make([][]map[string]string,0)
	for _, table := range data {
		// Extract table content
		content := table[0]
		
		// Split the content by newline to get individual rows
		rows := strings.Split(content, "\n")
		
		// Extract headers from the first row
		headers := strings.Split(rows[1],"\t")

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

		
	}
	if err := encoder.Encode(alltable); err != nil {
		log.Printf("Error writing JSON data: %v", err)
	}
	fmt.Printf("JSON file for %s created successfully .\n", filename)

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
