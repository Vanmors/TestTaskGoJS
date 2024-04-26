package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://hypeauditor.com/top-instagram-all-russia/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var dataArray []Data

	doc.Find(".row").Each(func(i int, s *goquery.Selection) {
		rank := s.Find(".row-cell.rank").Text()
		name := s.Find(".contributor__name-content").Text()
		category := s.Find(".row-cell.category").Text()
		subscribers := s.Find(".row-cell.subscribers").Text()
		audience := s.Find(".row-cell.audience").Text()
		authentic := s.Find(".row-cell.authentic").Text()
		engagement := s.Find(".row-cell.engagement").Text()

		if name != "" {
			data := Data{Rank: rank, Name: name, Category: category, Followers: subscribers, Country: audience, Authentic: authentic, Engagement: engagement}
			dataArray = append(dataArray, data)
		}
	})

	c := NewCsvWorker("data.csv")
	c.Write(dataArray)

}
