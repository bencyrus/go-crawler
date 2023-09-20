package main

import (
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func setupCollector() *colly.Collector {
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Failed to fetch the website: %v\n", err)
	})

	return c
}

func scrapeMainPage() []Company {
	var companies []Company
	mainCollector := setupCollector()

	queryTable(mainCollector, &companies)

	q, _ := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 10000})

	// Create a channel to sync
	syncChannel := make(chan bool)

	q.AddURL("https://www.crunchbase.com/hub/toronto-companies-fewer-than-1000-employees/hub_overview_default/top?tab=top_orgs")
	go func() {
		q.Run(mainCollector)
		syncChannel <- true // Send a true value into the channel when scraping is complete
	}()

	<-syncChannel // Wait for true value from channel
	return companies
}
