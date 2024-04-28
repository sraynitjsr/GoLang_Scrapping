package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	fmt.Print("****Scrapping Using GoLang****\n\n")

	fmt.Println("Parsing JSON Data")
	parseJSONData()
}

func parseJSONData() {
	data, err := os.ReadFile("DATA/json_data.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, p := range people {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
		fmt.Println("City:", p.City)
		fmt.Println()
	}
}
