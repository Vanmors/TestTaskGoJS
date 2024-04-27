package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type HtmlParser struct {
	path string
}

func NewHtmlParser(Path string) *HtmlParser {
	return &HtmlParser{path: Path}
}

func (h *HtmlParser) parse() []Data {
	res, err := http.Get(h.path)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

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

		category = insertAmpersand(category)

		if name != "" {
			data := Data{Rank: rank, Name: name, Category: category, Followers: subscribers, Country: audience, Authentic: authentic, Engagement: engagement}
			dataArray = append(dataArray, data)
		}
	})
	return dataArray
}

func insertAmpersand(str string) string {
	var sb strings.Builder
	for i, char := range str {
		if i > 0 && 'A' <= char && char <= 'Z' && str[i-1] != ' ' {
			sb.WriteRune(' ')
			sb.WriteRune('&')
			sb.WriteRune(' ')
		}
		sb.WriteRune(char)
	}
	return sb.String()
}
