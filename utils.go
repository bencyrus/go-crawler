package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func saveAsJSON(filename string, companies []Company) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	// Use json.MarshalIndent to pretty-print the JSON
	data, err := json.MarshalIndent(companies, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v\n", err)
	}

	// Write the formatted JSON data to the file
	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Failed to write to file: %v\n", err)
	}
}

func displayTable(companies []Company) {
	for _, company := range companies {
		fmt.Printf("Company Name: %s\n", company.CompanyName)
		fmt.Printf("Total Funding Amount: %s\n", company.TotalFundingAmount)
		fmt.Printf("CB Rank (Organization): %s\n", company.CBRankOrganization)
		fmt.Printf("Trend Score (30 Days): %s\n", company.TrendScore30Days)
		fmt.Printf("Company URL: %s\n", company.CompanyURL)
		fmt.Println("---------------")
	}
}

func displayTableWithDetails(companies []Company) {
	for _, company := range companies {
		fmt.Printf("Company Name: %s\n", company.CompanyName)
		fmt.Printf("Total Funding Amount: %s\n", company.TotalFundingAmount)
		fmt.Printf("CB Rank (Organization): %s\n", company.CBRankOrganization)
		fmt.Printf("Trend Score (30 Days): %s\n", company.TrendScore30Days)
		fmt.Printf("Company URL: %s\n", company.CompanyURL)
		fmt.Printf("Description: %s\n", company.Description)
		fmt.Printf("Number of Employees: %s\n", company.NumberOfEmployees)
		fmt.Printf("Last Funding Type: %s\n", company.LastFundingType)
		fmt.Printf("Website URL: %s\n", company.WebsiteURL)
		fmt.Printf("Rank: %s\n", company.Rank)
		fmt.Println("---------------")
	}
}

func getFirstElement(elements []string) string {
	if len(elements) > 0 {
		return elements[0]
	}
	return ""
}
