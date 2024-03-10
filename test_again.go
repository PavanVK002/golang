package main

import (
	"fmt"
	"io"
	"log"
	"time"
	"net/http"
	"os"
)

func main() {
	// URL for downloading the CSV file
	todayDate :=  time.Now().Format("02/03/2006") // Format
	url := fmt.Sprintf("https://api.bseindia.com/BseIndiaAPI/api/StockPriceCSVDownload/w?pageType=0&rbType=D&Scode=544083&FDates=01/01/2018&TDates=%s",todayDate)
	// url := "https://api.bseindia.com/BseIndiaAPI/api/StockPriceCSVDownload/w?pageType=0&rbType=D&Scode=500012&FDates=01/03/2018&TDates=01/03/2024"
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
	statusCode := response.StatusCode
	statusText := response.Status

	fmt.Printf("HTTP Status Code: %d\n", statusCode)
	fmt.Printf("HTTP Status Text: %s\n", statusText)
	// Create a new file to save the CSV
	file, err := os.Create("bse_data.csv")
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Failed to save CSV data: %v", err)
	}

	fmt.Println("CSV file created successfully.")
}
