package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/chromedp/chromedp"
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

func FetchAndSave(url, fileName string, client *http.Client, headers map[string]string) error {
	fmt.Printf("Starting to fetch data from: %s\n", url)

	// Create a new context with logging
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.UserAgent(headers["User-Agent"]),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Set a timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var htmlContent string
	
	// Run chromedp tasks
	err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(url),
		chromedp.WaitVisible(`.table .row[data-v-bf890aa6]`),
		// First get the HTML content to verify what we're working with
		chromedp.OuterHTML(`.table`, &htmlContent),
		chromedp.Sleep(2*time.Second), // Give the page a moment to fully render
	)
	if err != nil {
		return fmt.Errorf("failed to fetch page: %v", err)
	}

	fmt.Printf("Found table HTML: %s\n", htmlContent[:200]) // Print first 200 chars for debugging

	var rows []struct {
		Rank      string `json:"rank"`
		Nick      string `json:"nick"`
		FirstName string `json:"firstName"`
		Category  string `json:"category"`
		Followers string `json:"followers"`
		Country   string `json:"country"`
		EngAuth   string `json:"engAuth"`
		EngAvg    string `json:"engAvg"`
	}

	// Extract data using JavaScript
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll('.table .row[data-v-bf890aa6]')).map(row => {
				console.log('Processing row:', row.textContent); // Debug logging
				return {
					rank: row.querySelector('.row-cell.rank span[data-v-bf890aa6]')?.textContent?.trim() || '',
					nick: row.querySelector('.contributor__content-username')?.textContent?.trim() || '',
					firstName: row.querySelector('.contributor__content-fullname')?.textContent?.trim() || '',
					category: row.querySelector('.row-cell.category .tag__content')?.textContent?.trim() || '',
					followers: row.querySelector('.row-cell.subscribers')?.textContent?.trim() || '',
					country: row.querySelector('.row-cell.audience')?.textContent?.trim() || '',
					engAuth: row.querySelector('.row-cell.authentic')?.textContent?.trim() || '',
					engAvg: row.querySelector('.row-cell.engagement')?.textContent?.trim() || ''
				};
			})
		`, &rows),
	)
	if err != nil {
		return fmt.Errorf("failed to extract data: %v", err)
	}

	fmt.Printf("Extracted %d rows\n", len(rows))

	// Convert to our Row type
	var finalRows []Row
	for _, r := range rows {
		finalRows = append(finalRows, Row{
			rank:      r.Rank,
			nick:      r.Nick,
			firstName: r.FirstName,
			category:  r.Category,
			followers: r.Followers,
			country:   r.Country,
			engAuth:   r.EngAuth,
			engAvg:    r.EngAvg,
		})
	}

	if len(finalRows) == 0 {
		return fmt.Errorf("no data found in page")
	}

	// Create CSV file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func() {
		file.Close()
		fmt.Printf("Closed file: %s\n", fileName)
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	headers_csv := []string{"Rank", "Nickname", "Full Name", "Category", "Followers", "Country", "Engagement Auth", "Engagement Avg"}
	if err := writer.Write(headers_csv); err != nil {
		return fmt.Errorf("error writing headers: %v", err)
	}

	// Write data rows
	for _, row := range finalRows {
		// Process category string
		var categoryString strings.Builder
		for idx, val := range row.category {
			if idx > 0 {
				if unicode.IsUpper(val) && !unicode.IsUpper(rune(row.category[idx-1])) && unicode.IsLetter(val) {
					categoryString.WriteString(" ")
				}
			}
			categoryString.WriteRune(val)
		}

		record := []string{
			row.rank,
			row.nick,
			row.firstName,
			categoryString.String(),
			row.followers,
			row.country,
			row.engAuth,
			row.engAvg,
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("error writing row: %v", err)
		}
	}

	fmt.Printf("Successfully saved %d rows to %s\n", len(finalRows), fileName)
	return nil
}
