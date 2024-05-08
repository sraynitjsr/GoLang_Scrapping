package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	scrapedJSONData := scrapeJSONData("https://raw.githubusercontent.com/sraynitjsr/GoLang_Scrapping/main/Online_Data_Scrapping/data1.json")
	scrapedCSVData := scrapeCSVData("https://raw.githubusercontent.com/sraynitjsr/GoLang_Scrapping/main/CSV/csv_data.csv")

	fmt.Println("Scraped JSON Data:")
	for _, obj := range scrapedJSONData {
		fmt.Println(obj)
	}

	fmt.Println("\nScraped CSV Data:")
	for _, row := range scrapedCSVData {
		fmt.Println(row)
	}
}

func scrapeJSONData(url string) []interface{} {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer resp.Body.Close()

	var result []interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	return result
}

func scrapeCSVData(url string) [][]string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}

	return records
}
