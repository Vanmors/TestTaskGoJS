package main

func main() {
	h := NewHtmlParser("https://hypeauditor.com/top-instagram-all-russia/")
	c := NewCsvWorker("data.csv")

	dataArray := h.parse()
	c.Write(dataArray)

}
