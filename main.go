package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

type Row struct {
	rank      string
	nick      string
	firstName string
	category  string
	followers string
	country   string
	engAuth   string
	engAvg    string
}

func FetchAndSave(url, fileName string) {
	var Rows []Row

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/111.0.0.0 Safari/537.36"),
		colly.Async(true), // Enable asynchronous requests
	)

	// Set a longer timeout
	c.SetRequestTimeout(60 * time.Second)

	// Limit the number of concurrent requests
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       5 * time.Second, // Add a delay between requests
	})

	c.OnHTML(".table .row[data-v-bf890aa6]", func(element *colly.HTMLElement) {
		rank := element.ChildText(".row-cell.rank span[data-v-bf890aa6]")
		nick := element.ChildText(".contributor__content-username")
		firstName := element.ChildText(".contributor__content-fullname")
		category := element.ChildText(".row-cell.category .tag__content")
		followers := element.ChildText(".row-cell.subscribers")
		country := element.ChildText(".row-cell.audience")
		engAuth := element.ChildText(".row-cell.authentic")
		engAvg := element.ChildText(".row-cell.engagement")

		var categoryString strings.Builder
		for idx, val := range category {
			if idx > 0 {
				if unicode.IsUpper(val) && !unicode.IsUpper(rune(category[idx-1])) && unicode.IsLetter(val) {
					categoryString.WriteString(" ")
				}
			}
			categoryString.WriteRune(val)
		}

		row := Row{
			rank:      rank,
			nick:      nick,
			firstName: firstName,
			category:  categoryString.String(),
			followers: followers,
			country:   country,
			engAuth:   engAuth,
			engAvg:    engAvg,
		}

		Rows = append(Rows, row)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with response: %v\nError: %v", r.Request.URL, r, err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	c.Wait() // Wait for all asynchronous tasks to complete

	if len(Rows) == 0 {
		log.Println("No data found. Please check the selectors and the target URL.")
		return
	}

	csvFile, csvErr := os.Create(fileName)
	if csvErr != nil {
		log.Fatalln("Failed to create the output CSV file", csvErr)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	headers := []string{
		"rank",
		"nick",
		"Name",
		"category",
		"followers",
		"country",
		"engAuth",
		"engAvg",
	}
	writer.Write(headers)

	for _, row := range Rows {
		record := []string{
			row.rank,
			row.nick,
			row.firstName,
			row.category,
			row.followers,
			row.country,
			row.engAuth,
			row.engAvg,
		}
		writer.Write(record)
	}
}
