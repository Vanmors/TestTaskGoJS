package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	apiURL         = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
	updateInterval = 10 * time.Minute
)

func main() {

	c := NewCoinWorker(apiURL, updateInterval)
	var coins []CoinData
	var mutex sync.Mutex
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите название: ")

	scanner.Scan()
	symbol := scanner.Text()
	symbol = strings.TrimSpace(symbol)

	wg.Add(1)
	go c.UpdateCoinData(apiURL, &coins, &mutex, &wg)
	wg.Wait()

	for {
		mutex.Lock()
		price, err := c.GetCurrencyPrice(coins, symbol)
		mutex.Unlock()

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Price of %s: %.2f USD\n", symbol, price)
		}

		time.Sleep(1 * time.Minute) // Проверка курса каждую минуту
	}
}
