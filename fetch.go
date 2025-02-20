package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"os"
	"time"
	"path/filepath"
)

func FetchAndSave(url string, fileName string, client *http.Client, headers map[string]string) error {
	fmt.Printf("Creating file: %s\n", fileName)
	
	// 确保文件名是正确的路径格式
	fileName = filepath.Clean(fileName)
	
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

	c := colly.NewCollector(
		colly.UserAgent(headers["User-Agent"]),
		colly.MaxDepth(1),
		colly.AllowURLRevisit(),
		colly.DisallowedDomains("www.google-analytics.com", "googleads.g.doubleclick.net"),
	)

	c.SetRequestTimeout(30 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		for key, value := range headers {
			r.Headers.Set(key, value)
		}
		fmt.Printf("Making request to: %s\n", r.URL.String())
	})

	rowCount := 0

	c.OnHTML("div.table div.row", func(e *colly.HTMLElement) {
		row := []string{
			e.ChildText("div.rank"),
			e.ChildText("div.nick"),
			e.ChildText("div.Name"),
			e.ChildText("div.category"),
			e.ChildText("div.followers"),
			e.ChildText("div.country"),
			e.ChildText("div.engAuth"),
			e.ChildText("div.engAvg"),
		}

		// 检查行是否为空
		hasData := false
		for _, field := range row {
			if field != "" {
				hasData = true
				break
			}
		}

		if !hasData {
			fmt.Printf("Skipping empty row\n")
			return
		}

		if err := writer.Write(row); err != nil {
			fmt.Printf("Error writing row: %v\n", err)
			return
		}
		rowCount++
		if rowCount%10 == 0 {
			fmt.Printf("Processed %d rows...\n", rowCount)
			writer.Flush() // 定期刷新写入
		}
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Visited URL: %s (Status: %d)\n", r.Request.URL, r.StatusCode)
		if r.StatusCode != 200 {
			fmt.Printf("Response body: %s\n", string(r.Body))
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error visiting %s: %v\n", r.Request.URL, err)
		if r != nil {
			fmt.Printf("Response body: %s\n", string(r.Body))
		}
	})

	fmt.Printf("Starting to scrape URL: %s\n", url)
	err = c.Visit(url)
	if err != nil {
		return fmt.Errorf("error visiting URL: %v", err)
	}

	if rowCount == 0 {
		return fmt.Errorf("no data found for URL: %s", url)
	}

	fmt.Printf("Completed scraping. Total rows processed: %d\n", rowCount)
	return nil
}
