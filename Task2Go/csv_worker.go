package main

import (
	"encoding/csv"
	"log"
	"os"
)

type CsvWorker struct {
	path string
}

func NewCsvWorker(Path string) *CsvWorker {
	return &CsvWorker{path: Path}
}

func (c *CsvWorker) Write(data []Data) {
	file, err := os.Create(c.path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Rank", "Name", "Category", "Followers", "Country", "Authentic", "Engagement"}
	if err := writer.Write(headers); err != nil {
		log.Fatal(err)
	}

	for _, d := range data {
		record := []string{
			d.Rank,
			d.Name,
			d.Category,
			d.Followers,
			d.Country,
			d.Authentic,
			d.Engagement,
		}
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Данные успешно записаны в файл data.csv")
}

func arrayToString(arr []string) string {
	str := ""
	for i, v := range arr {
		str += v
		if i < len(arr)-1 {
			str += ", "
		}
	}
	return str
}
