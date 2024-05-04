package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://raw.githubusercontent.com/sraynitjsr/GoLang_Scrapping/main/Online_Data_Scrapping/data1.json"
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

	for _, obj := range result {
		fmt.Println(obj)
	}
}
