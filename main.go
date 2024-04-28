package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
	City string `json:"city" yaml:"city"`
}

func main() {
	fmt.Print("****Scrapping Using GoLang****\n\n")

	parseJSONData()

	parseYAMLData()

	parseCSVData()
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

	fmt.Println("Parsed JSON data:")
	for _, p := range people {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
		fmt.Println("City:", p.City)
		fmt.Println()
	}
}

func parseYAMLData() {
	data, err := os.ReadFile("DATA/yaml_data.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	var people []Person
	err = yaml.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return
	}

	fmt.Println("Parsed YAML data:")
	for _, p := range people {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
		fmt.Println("City:", p.City)
		fmt.Println()
	}
}

func parseCSVData() {
	file, err := os.Open("DATA/csv_data.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	var people []Person
	for _, row := range records {
		age, _ := strconv.Atoi(row[1])
		person := Person{
			Name: row[0],
			Age:  age,
			City: row[2],
		}
		people = append(people, person)
	}

	fmt.Println("Parsed CSV data:")
	for _, p := range people {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
		fmt.Println("City:", p.City)
		fmt.Println()
	}
}
