package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
	City string `json:"city" yaml:"city"`
}

func main() {
	fmt.Print("****Scrapping Using GoLang****\n\n")

	fmt.Println("Parsing JSON Data")
	parseJSONData()

	fmt.Println("\nParsing YAML Data")
	parseYAMLData()
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

	for _, p := range people {
		fmt.Println("Name:", p.Name)
		fmt.Println("Age:", p.Age)
		fmt.Println("City:", p.City)
		fmt.Println()
	}
}
