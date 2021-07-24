package main

import (
	"fmt"
	"github.com/Badmullafo/alphavantage/golang_web/pkg/request"
	"log"
	"net/http"
)

func main() {

	apiKey, symbol := "RABZYXWVHB8MX5GO", "IBM"
	url := "https://www.alphavantage.co/query?apikey=" + apiKey + "&function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + symbol

	doget := get.Get(apiKey, symbol, url)

	fmt.Printf("%+v\n", doget)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
