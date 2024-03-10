package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	client = &http.Client{}
	logger *log.Logger
)

func main() {
	logFile, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v\n", err)
	}
	defer logFile.Close()
	logger = log.New(logFile, "", log.LstdFlags)

	channel := make(chan string)
	fmt.Printf("Start time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))
	logger.Printf("Start time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

	inputFile, err := os.Open("Security_codes.csv")
	if err != nil {
		logger.Fatalf("Error opening input file: %v\n", err)
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)

	records, err := reader.ReadAll()
	if err != nil {
		logger.Fatalf("Error reading input CSV: %v\n", err)
	}

	for i := 1; i < 20; i++ {
		go Webscrapping_IndexHighlight(channel)
	}

	for _, record := range records {
		channel <- record[0]
	}

	fmt.Printf("End time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))
	logger.Printf("End time of execution: %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

}

func Webscrapping_IndexHighlight(channel chan string) {
	for sc := range channel {
		todayDate := time.Now().Format("02/03/2006")
		url := fmt.Sprintf("https://api.bseindia.com/BseIndiaAPI/api/StockPriceCSVDownload/w?pageType=0&rbType=D&Scode=%s&FDates=01/01/2018&TDates=%s", sc, todayDate)

		filename := fmt.Sprintf("Securities_equity_%s.csv", sc)

		const maxRetries = 4
		var body []byte
		var availabilityCheck bool
		availabilityCheck = false
		var statusCode int
		var statusText string
		var retries int
		for retries = 0; retries < maxRetries; retries++ {
			response, err := createHTTPRequest(url)
			if err != nil {
				logger.Printf("Failed to send request on %s: %v\n", time.Now().Format("2006-01-02 15:04:05.000"), err)
				continue
			}
			defer response.Body.Close()
			statusCode = response.StatusCode
			statusText = response.Status
			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				logger.Printf("Failed to copy response body on %s: %v\n", time.Now().Format("2006-01-02 15:04:05.000"), err)
				continue
			}

			bodyStr := string(body)
			if len(bodyStr) == 0 {
				availabilityCheck = true
				logger.Printf("CSV Content for %s is empty on %s for retry %d\n", sc, time.Now().Format("2006-01-02 15:04:05.000"), retries+1)
				continue
			}
			if statusCode == 503 {
				availabilityCheck = true
				if retries == 3 {
					logger.Printf("Failed to extract CSV Content for %s due to %s on %s for retry %d\n", sc, statusText, time.Now().Format("2006-01-02 15:04:05.000"), retries+1)
				}
				continue
			}

			availabilityCheck = false
			break
		}

		if availabilityCheck != true {
			saveCSV(body, filename, retries)
		}else if statusCode == 503 {
			logger.Printf("CSV file for %s was not created due to %s on %s with %d retries.\n", filename,statusText, time.Now().Format("2006-01-02 15:04:05.000"), retries)

		}else{
			logger.Printf("CSV file for %s was not created due to reponse being empty on %s with %d retries.\n", filename, time.Now().Format("2006-01-02 15:04:05.000"), retries)
		}
	}
}

func createHTTPRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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

	time.Sleep(5 * time.Second)
	response, err := client.Do(req)

	return response, err
}

func saveCSV(body []byte, filename string, retries int) {
	file, err := os.Create("security/" + filename)
	if err != nil {
		logger.Printf("Failed to create CSV file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		logger.Printf("Failed to save CSV data: %v\n", err)
		return
	}
	logger.Printf("CSV file for %s created successfully on %s with %d retries.\n", filename, time.Now().Format("2006-01-02 15:04:05.000"), retries)
}
