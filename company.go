package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Company struct {
	CompanyName        string `json:"company_name"`
	TotalFundingAmount string `json:"total_funding_amount"`
	CBRankOrganization string `json:"cb_rank_organization"`
	TrendScore30Days   string `json:"trend_score_30_days"`
	CompanyURL         string `json:"company_url"`
	Description        string `json:"description,omitempty"`
	NumberOfEmployees  string `json:"number_of_employees,omitempty"`
	LastFundingType    string `json:"last_funding_type,omitempty"`
	WebsiteURL         string `json:"website_url,omitempty"`
	Rank               string `json:"rank,omitempty"`
}

func queryTable(c *colly.Collector, companies *[]Company) {
	c.OnHTML("table.card-grid", func(e *colly.HTMLElement) {
		firstTh := e.ChildText("thead tr th:first-child")
		if strings.Contains(firstTh, "Organization Name") {
			e.ForEach("tbody", func(_ int, tbodyHtml *colly.HTMLElement) {
				extractCompanyDetails(tbodyHtml, companies)
			})
		}
	})
}

func extractCompanyDetails(tbodyHtml *colly.HTMLElement, companies *[]Company) {
	tbodyHtml.ForEach("tr", func(_ int, rowHtml *colly.HTMLElement) {
		var columns []string
		var companyURL string

		rowHtml.ForEach("td", func(index int, cell *colly.HTMLElement) {
			if index == 0 {
				companyURL = "https://www.crunchbase.com/" + cell.ChildAttr("a", "href")
			}
			columns = append(columns, strings.TrimSpace(cell.Text))
		})

		if len(columns) > 0 {
			*companies = append(*companies, Company{
				CompanyName:        columns[0],
				TotalFundingAmount: columns[1],
				CBRankOrganization: columns[2],
				TrendScore30Days:   columns[3],
				CompanyURL:         companyURL,
			})
		}
	})
}

func scrapeAdditionalDetails(company *Company) {
	c := setupCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Fill in additional details here
		company.Description = getFirstElement(e.ChildTexts("span.description"))
		company.NumberOfEmployees = getFirstElement(e.ChildTexts(`a[href*="num_employees_enum"]`))
		company.LastFundingType = getFirstElement(e.ChildTexts(`a[href*="last_funding_type"]`))
		company.WebsiteURL = e.ChildAttr("a[rel='nofollow noopener noreferrer']", "href")
		company.Rank = getFirstElement(e.ChildTexts(`a[href*="rank_org_company"]`))
	})

	err := c.Visit(company.CompanyURL)
	if err != nil {
		log.Printf("Failed to visit company URL: %v\n", err)
	}
}
