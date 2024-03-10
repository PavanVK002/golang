package main

import (
	"fmt"
	"encoding/csv"
	// "io"
	"io/ioutil"
	"log"
	"time"
	
	// "strings"
	"net/http"
	"os"
)
var (
	client = &http.Client{}
)

func main() {
	channel := make(chan string)
	fmt.Printf("Start time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

	inputFile, err := os.Open("Security_codes_1000.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v \n", err)
	}
	defer inputFile.Close()

	// Create a CSV reader for the input file
	reader := csv.NewReader(inputFile)

	// Read all records from the input CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading input CSV: %v \n", err)
	}
	
	for i :=1 ;i<20;i++{
		go Webscrapping_IndexHighlight(channel)
	}

	for _,record := range records{
		channel <- record[0]
	}
	
	// Webscrapping_IndexHighlight("500012")
	// Wait for goroutines to finish
	// time.Sleep(10 * time.Minute)
	fmt.Printf("End time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

	
}

func Webscrapping_IndexHighlight(channel chan string){
	// defer wg.Done()
	for sc := range channel{
		// time.Sleep(2*time.Second) // Avoid being blocked by
		todayDate :=  time.Now().Format("02/03/2006") // Format
	url := fmt.Sprintf("https://api.bseindia.com/BseIndiaAPI/api/StockPriceCSVDownload/w?pageType=0&rbType=D&Scode=%s&FDates=01/01/2018&TDates=%s",sc,todayDate)

	 //"https://api.bseindia.com/BseIndiaAPI/api/MktCapDownloadCSV/w?type=2&cat=%d "the url is for accessing the csv
	//https://api.bseindia.com/BseIndiaAPI/api/MktCapBoard/w?cat=%d&type=2 url is for getting data in json
			// change the cat value to get diffrent categories in order ,do not change type


	filename := fmt.Sprintf("Securities_equity_%s.csv",sc)
	// Create a new HTTP client with custom headers
	// client := &http.Client{}

	const maxRetries = 5 // Maximum number of retries
	var body []byte
	var availability_check bool
	availability_check = false
	var  retries int 
	for retries = 0; retries < maxRetries; retries++ {
		// 
		response, err := create_http_request(url)
		if err != nil {
			fmt.Printf("Failed to send request on %s: %v \n", time.Now().Format("2006-01-02 15:04:05.000"), err)
			continue // Retry the request
		}
		defer response.Body.Close()
		statusCode := response.StatusCode
		statusText := response.Status
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Failed to copy response body on %s: %v \n", time.Now().Format("2006-01-02 15:04:05.000"), err)
			continue // Retry the request
		}
	
		// Convert the response body to a string
		bodyStr := string(body)
		if len(bodyStr)==0{
			availability_check = true
			// Service unavailable, retry after backoff duration
			fmt.Printf(" CSV Content for %s is empty on %s for retry %d \n", sc,  time.Now().Format("2006-01-02 15:04:05.000"), retries+1)
			// time.Sleep(3*time.Second)
			continue
		}
		// Check if the response body contains "503 - Service Not Available"
		if statusCode == 503  {
			availability_check = true
			if retries==4{
				fmt.Printf("Failed to extract CSV Content for %s due to %s on %s for retry %d \n", sc,statusText, time.Now().Format("2006-01-02 15:04:05.000"),retries+1)
			}
			// time.Sleep(3*time.Second)
			continue // Retry the request
		}
	
		// If we reach this point, the request was successful, so break out of the loop
		availability_check = false
		break
	}

	// Create a new file to save the CSV
	if availability_check != true{
		CSV_Save(body ,filename,retries)
	}
	 
}
}


func create_http_request(url string) ( *http.Response, error) {
	// Create a new request with the URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v \n", err)
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


	time.Sleep(6*time.Second)
	response, err := client.Do(req)
	
	return response ,err
}	 


func CSV_Save(body [] byte,filename string,retries int){
	file, err := os.Create("security/"+filename)
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v \n", err)
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = file.Write(body)
	if err != nil {
			log.Fatalf("Failed to save CSV data: %v \n", err)
	}
	fmt.Printf("CSV file for %s created successfully on %s with %d retries.\n",filename,time.Now().Format("2006-01-02 15:04:05.000"),retries)
	
}
	




