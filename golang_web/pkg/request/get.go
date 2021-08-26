package request

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	layoutISO = "2006-01-02"
	tz        = "US/Eastern"
	layoutUS  = "January 2, 2006"
	urlS      = "https://www.alphavantage.co/query?apikey="
	rtype     = "&function=TIME_SERIES_DAILY_ADJUSTED&symbol="
)

var total float64

func GetJson(apiKey, symbol string, nDays int) (float64, error) {

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
		return 0.0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return 0.0, err
	}

	tsd := gjson.GetBytes(body, "Time Series (Daily)")

	//Outer loop
	tsd.ForEach(func(key, value gjson.Result) bool {

		date, err := time.Parse(layoutISO, key.String())
		// fmt.Println(date, err)

		if err != nil {
			fmt.Println(err)
		}

		diff := currentTime.Sub(date).Hours()

		//Convert days to hours
		daysH := float64(nDays) * 24

		if diff < daysH {

			//In hours
			//fmt.Printf("Hours: %f\n", diff)
			//	fmt.Printf("Hours: %f\n", days)

			//println(value.String())
			//Inner loop
			value.ForEach(func(key, value gjson.Result) bool {

				if k := key.String(); k == "4. close" {
					v := value.Float()
					fmt.Printf("The key is:%s, the value is:%f\n", k, v)
					total = total + v
				}
				return true // keep iterating
			})
		}
		return true // keep iterating
	})
	return total, nil
}
