package main

import (
	"fmt"
	// "encoding/csv"
	// "io"
	"io/ioutil"
	"log"
	"time"
	// "sync"
	"strings"
	"net/http"
	"os"
)

// var wg sync.WaitGroup
func main() {

	fmt.Printf("Start time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

	// inputFile, err := os.Open("Security_codes.csv")
	// if err != nil {
	// 	log.Fatalf("Error opening input file: %v \n", err)
	// }
	// defer inputFile.Close()

	// // Create a CSV reader for the input file
	// reader := csv.NewReader(inputFile)

	// // Read all records from the input CSV
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	log.Fatalf("Error reading input CSV: %v \n", err)
	// }
	// wg.Add(len(records))
	// for _,record := range records{
	// 	go Webscrapping_IndexHighlight(record[0])
	// }
	// wg.Wait()
	Webscrapping_IndexHighlight("500012")
	// Wait for goroutines to finish
	// time.Sleep(10 * time.Second)
	fmt.Printf("End time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

	
}
func Webscrapping_IndexHighlight(security_code string){
	// defer wg.Done()
	todayDate :=  time.Now().Format("02/03/2006") // Format
	fmt.Println(todayDate)
	url := fmt.Sprintf("https://api.bseindia.com/BseIndiaAPI/api/StockPriceCSVDownload/w?pageType=0&rbType=D&Scode=%s&FDates=01/03/2018&TDates=%s",security_code,todayDate)
	// url := 				  "https://api.bseindia.com/BseIndiaAPI/api/StockPriceCSVDownload/w?pageType=0&rbType=D&Scode=500012&FDates=01/03/2018&TDates=02/03/2024"
	 //"https://api.bseindia.com/BseIndiaAPI/api/MktCapDownloadCSV/w?type=2&cat=%d "the url is for accessing the csv
	//https://api.bseindia.com/BseIndiaAPI/api/MktCapBoard/w?cat=%d&type=2 url is for getting data in json
			// change the cat value to get diffrent categories in order ,do not change type

	fmt.Println(url)
	filename := fmt.Sprintf("Securities_equity_%s.csv",security_code)
	// Create a new HTTP client with custom headers
	client := &http.Client{}


	// client := &http.Client{
	// 	Timeout: 2 * time.Second, // Set a timeout for the HTTP request
	// }
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

	// Send the request
	response, err := client.Do(req)
	if err != nil  {
		fmt.Printf("Failed to send request on %s: %v \n",time.Now().Format("2006-01-02 15:04:05.000"), err)
		return
	}
	defer response.Body.Close()
	// fmt.Println(response)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to copy response body on %s: %v \n", time.Now().Format("2006-01-02 15:04:05.000"), err)
		return
	}
	// fmt.Println(body)
	// Convert the response body to a string
	bodyStr := string(body)

	// Check if the response body contains "503 - Service Not Available"
	if strings.Contains(bodyStr, "503 - Service Not Available") {
		fmt.Printf("Failed to extract CSV Content for %s on %s \n",security_code, time.Now().Format("2006-01-02 15:04:05.000"))
		return
	}

	// Create a new file to save the CSV
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
	fmt.Printf("CSV file for %s created successfully on %s .\n",filename,time.Now().Format("2006-01-02 15:04:05.000"))
	
}


