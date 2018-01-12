package lib

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
)

var apiKey = "your_alphavantage_api_key"

// MSFT = Microsoft, GOOGL = Alphabet Inc., TSLA = Tesla, AMZN = Amazon, DIS = Disney
var stockSymbols = []string{
	"GOOGL",
	"TSLA",
	"AAPL",
	"AMZN",
	"DIS",
	}

// In-memory hashtable in which N amount of stocks will "live in"
var MarketDataMap = make(map[string]interface{})

// Export GetStockMarketData
func GetStockData() {

	for _, stockSymbol := range stockSymbols {
		var requestLink = fmt.Sprintf(
			"https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&symbol=%s&apikey=%s",
			stockSymbol, apiKey)

		response, err := http.Get(requestLink)

		fmt.Println("Getting data for.. " + stockSymbol)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			rawData, _ := ioutil.ReadAll(response.Body)

			data := make(map[string]interface{})

			err := json.Unmarshal(rawData, &data)

			if err != nil {
				fmt.Print(err)
				return
			}
			MarketDataMap[stockSymbol] = &data
		}
	}
	// Recursive call every n seconds
	time.AfterFunc(time.Second * 10, GetStockData)
}
