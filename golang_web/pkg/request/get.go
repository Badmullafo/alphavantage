package request

import (
	//   "bytes"

	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	_ "log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

const (
	layoutISO = "2006-01-02"
	tz        = "US/Eastern"
	layoutUS  = "January 2, 2006"
	urlS      = "https://www.alphavantage.co/query?apikey="
	rtype     = "&function=TIME_SERIES_DAILY_ADJUSTED&symbol="
)

func GetJson(apiKey, symbol string) error {

	//apiKey, symbol := "RABZYXWVHB8MX5GO", "IBM"
	url := urlS + apiKey + rtype + symbol
	//init the loc
	loc, _ := time.LoadLocation(tz)

	//set timezone,
	currentTime := time.Now().In(loc)
	fmt.Println("Geting data from", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	md := gjson.GetBytes(body, "Time Series (Daily)")

	if err != nil {
		log.Fatal(err)
		return err
	}

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
	return nil
}
