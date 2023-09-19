package main

func main() {
	companies := scrapeMainPage()

	// Display the initial data
	displayTable(companies)
	// Save the initial data as JSON
	saveAsJSON("data/companies.json", companies)

	// Loop through each company to get additional details
	for i := range companies {
		scrapeAdditionalDetails(&companies[i])
	}

	// Save and display the data with additional details
	saveAsJSON("data/companies-with-details.json", companies)
	displayTableWithDetails(companies)
}
