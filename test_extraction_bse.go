package main

import (
	"fmt"
	"io"
	"log"
	"time"
	"sync"
	"net/http"
	"os"
)
var category = []string{"Market_Cap_Broad", "Sector&Industry", "Thematics", "Startegy", "Sustainability", "Volatility", "Composite", "Government", "Corporate", "Money_Market"}

var wg sync.WaitGroup
func main() {

	fmt.Printf("Start time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go Webscrapping_IndexHighlight(i)
	}
	wg.Wait()

	// Wait for goroutines to finish
	// time.Sleep(10 * time.Second)
	fmt.Printf("End time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

	
}
func Webscrapping_IndexHighlight(type_ int){
	url := fmt.Sprintf("https://api.bseindia.com/BseIndiaAPI/api/MktCapBoard/w?cat=%d&type=2",type_)

	 //"https://api.bseindia.com/BseIndiaAPI/api/MktCapDownloadCSV/w?type=2&cat=%d "the url is for accessing the csv
	//https://api.bseindia.com/BseIndiaAPI/api/MktCapBoard/w?cat=%d&type=2 url is for getting data in json
			// change the cat value to get diffrent categories in order ,do not change type


	filename := fmt.Sprintf("BSE_Data_%s.json",category[type_-1])
	// Create a new HTTP client with custom headers
	client := &http.Client{}

	// Create a new request with the URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set custom headers
	req.Header.Set("authority", "api.bseindia.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("cookie", "_ga=GA1.1.570261728.1706230967; __gads=ID=a91014ba4ce71ca1:T=1706230966:RT=1706268552:S=ALNI_MaLyoAkOQBV9aRAh6HLuhrLITFgfg; _ga_TM52BJH9HF=GS1.1.1706268553.5.1.1706268553.0.0.0;        RT=\"z=1&dm=bseindia.com&si=319ac784-11cf-44d3-a8a5-8822b0723a43&ss=lruk8ovg&sl=1&tt=lz&rl=1&ld=100&ul=dls\"")
	req.Header.Set("referer", "https://www.bseindia.com/")
	req.Header.Set("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Microsoft Edge\";v=\"120\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()

	// Create a new file to save the CSV
	file, err := os.Create("csv_folder/"+filename)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Failed to save CSV data: %v", err)
	}

	fmt.Printf("CSV file for %s created successfully.\n",filename)
	wg.Done()
}


