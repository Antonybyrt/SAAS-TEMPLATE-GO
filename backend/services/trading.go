package services

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	krakenBaseURL = "https://api.kraken.com/0/public"
)

func fetchPairData(symbol string, ch chan<- models.TradingPair) {
	krakenSymbol := convertSymbol(symbol)

	url := fmt.Sprintf("%s/Ticker?pair=%s", krakenBaseURL, krakenSymbol)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", symbol, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response for %s: %v\n", symbol, err)
		return
	}

	var krakenResp models.KrakenResponse
	if err := json.Unmarshal(body, &krakenResp); err != nil {
		fmt.Printf("Error unmarshaling response for %s: %v\n", symbol, err)
		return
	}

	if len(krakenResp.Error) > 0 {
		fmt.Printf("Kraken API error for %s: %v\n", symbol, krakenResp.Error)
		return
	}

	var ticker models.KrakenTicker
	for _, t := range krakenResp.Result {
		ticker = t
		break
	}

	price, _ := strconv.ParseFloat(ticker.C[0], 64)

	openPrice, _ := strconv.ParseFloat(ticker.O, 64)
	change24h := ((price - openPrice) / openPrice) * 100

	pair := models.TradingPair{
		Symbol:    symbol,
		Price:     price,
		Change24h: change24h,
	}

	ch <- pair
}

func convertSymbol(symbol string) string {
	// Je mets les symboles kraken dans des symboles qu'on connait
	replacements := map[string]string{
		"BTC":  "XBT",
		"ETH":  "ETH",
		"SOL":  "SOL",
		"XTZ":  "XTZ",
		"USDC": "USDC",
	}

	parts := strings.Split(symbol, "/")
	if len(parts) != 2 {
		return symbol
	}

	base := replacements[parts[0]]
	quote := replacements[parts[1]]

	return base + quote
}

func GetTradingPairs() []models.TradingPair {
	pairs := []string{"BTC/USDC", "ETH/USDC", "SOL/USDC", "XTZ/USDC"}
	ch := make(chan models.TradingPair, len(pairs))
	var results []models.TradingPair

	for _, pair := range pairs {
		go fetchPairData(pair, ch)
	}

	timeout := time.After(5 * time.Second)
	for i := 0; i < len(pairs); i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Timeout waiting for results")
			return results
		}
	}

	return results
}
