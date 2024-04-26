package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type CoinWorker struct {
	apiURL         string
	updateInterval time.Duration
}

func NewCoinWorker(ApiURL string, UpdateInterval time.Duration) *CoinWorker {
	return &CoinWorker{apiURL: ApiURL, updateInterval: UpdateInterval}
}

func (c *CoinWorker) FetchCoinData(url string, wg *sync.WaitGroup) ([]CoinData, error) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var coins []CoinData
	err = json.Unmarshal(body, &coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}

func (c *CoinWorker) GetCurrencyPrice(coins []CoinData, symbol string) (float64, error) {
	for _, coin := range coins {
		if coin.Symbol == symbol {
			return coin.Price, nil
		}
	}
	return 0, fmt.Errorf("currency not found")
}

func (c *CoinWorker) UpdateCoinData(url string, coins *[]CoinData, mutex *sync.Mutex, wg *sync.WaitGroup) {
	for {
		newCoins, err := c.FetchCoinData(url, wg)
		if err != nil {
			fmt.Println("Error fetching coin data:", err)
			continue
		}

		mutex.Lock()
		*coins = newCoins
		mutex.Unlock()

		time.Sleep(updateInterval)
	}
}
