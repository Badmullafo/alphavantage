package main

import (
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/Badmullafo/alphavantage/golang_web/pkg/request"
	"github.com/tidwall/gjson"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func main() {

	apiKey, symbol := "RABZYXWVHB8MX5GO", "IBM"
	url := "https://www.alphavantage.co/query?apikey=" + apiKey + "&function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + symbol
	//init the loc
	loc, _ := time.LoadLocation("US/Eastern")

	//set timezone,
	currentTime := time.Now().In(loc)
	fmt.Println("Geting data from", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	md := gjson.GetBytes(body, "Time Series (Daily)")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting server at port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/

	md.ForEach(func(key, value gjson.Result) bool {

		date, err := time.Parse(layoutISO, key.String())
		// fmt.Println(date, err)

		if err != nil {
			fmt.Println(err)
		}

		diff := currentTime.Sub(date).Hours()

		if diff < 48.0 {

			//In hours
			fmt.Printf("Hours: %f\n", diff)
			//	fmt.Printf("Hours: %f\n", days)

			println(value.String())
		}

		return true // keep iterating
	})
}
