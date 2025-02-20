package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
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

func main() {
	// init slice structs
	var Rows []Row

	c := colly.NewCollector()

	// valid user-agent
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/111.0.0.0 Safari/537.36"

	//HTML elements
	c.OnHTML(".table .row[data-v-bf890aa6]", func(element *colly.HTMLElement) {
		rank := element.ChildText(".row-cell.rank span[data-v-bf890aa6]")
		nick := element.ChildText(".contributor__content-username")
		firstName := element.ChildText(".contributor__content-fullname")
		category := element.ChildText(".row-cell.category .tag__content")
		followers := element.ChildText(".row-cell.subscribers")
		country := element.ChildText(".row-cell.audience")
		engAuth := element.ChildText(".row-cell.authentic")
		engAvg := element.ChildText(".row-cell.engagement")

		// add space between categories
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

	// 错误处理
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// 开始抓取
	err := c.Visit("https://hypeauditor.com/top-instagram-clothing-outfits-united-states/")
	if err != nil {
		log.Fatal(err)
	}

	// 检查是否抓取到数据
	if len(Rows) == 0 {
		log.Println("No data found. Please check the selectors and the target URL.")
		return
	}

	// --- export to CSV ---

	// open the output CSV file
	csvFile, csvErr := os.Create("Instagram.csv")
	// if the file creation fails
	if csvErr != nil {
		log.Fatalln("Failed to create the output CSV file", csvErr)
	}
	// release the resource allocated to handle
	// the file before ending the execution
	defer csvFile.Close()

	// create a CSV file writer
	writer := csv.NewWriter(csvFile)
	// release the resources associated with the
	// file writer before ending the execution
	defer writer.Flush()

	// add the header row to the CSV
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

	// store each Industry product in the
	// output CSV file
	for _, row := range Rows {
		// convert the Industry instance to
		// a slice of strings
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

		// add a new CSV record
		writer.Write(record)
	}
}
