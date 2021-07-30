package main

import (
	"encoding/json"
	"fmt"
	"log"
	_ "net/http"

	"github.com/Badmullafo/alphavantage/golang_web/pkg/request"
)

func main() {

	apiKey, symbol := "RABZYXWVHB8MX5GO", "IBM"
	url := "https://www.alphavantage.co/query?apikey=" + apiKey + "&function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + symbol

	doget, err := request.Get(apiKey, symbol, url)

	if err != nil {
		log.Fatal(err)
	}

	prettyJSON, err := json.MarshalIndent(doget, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))

	fmt.Printf("Starting server at port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/
}
