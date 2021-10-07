package request

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
	timeout   = 1000
)

type result struct {
	symbol  string
	ndays   int
	dtype   string
	dateval map[time.Time]float64
}

func GetJson(apiKey, symbol string, nDays int) (result, error) {

	//apiKey, symbol := "RABZYXWVHB8MX5GO", "IBM"
	url := urlS + apiKey + rtype + symbol
	//init the loc
	loc, _ := time.LoadLocation(tz)

	//set timezone,
	currentTime := time.Now().In(loc)
	fmt.Println("Geting data from", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*timeout))
	defer cancel()
	req = req.WithContext(ctx)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	getRes := result{symbol, nDays, "", make(map[time.Time]float64)}
	// If there is an error message return here
	if err := gjson.GetBytes(out, "Error Message"); err.String() != "" {
		return result{}, errors.New(err.String())
	}

	tsd := gjson.GetBytes(out, "Time Series (Daily)")

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

			//Inner loop
			value.ForEach(func(key, value gjson.Result) bool {

				if k := key.String(); k == "4. close" {
					v := value.Float()
					//fmt.Printf("The key is:%s, the value is:%f\n", k, v)
					getRes.dateval[date] = v
				}
				return true // keep iterating
			})
		}
		return true // keep iterating
	})

	return getRes, nil
}

func (r *result) Getot() (float64, error) {

	var total float64

	for _, value := range r.dateval {
		total = total + value
	}
	return total, nil
}

func (r *result) Getavg() (float64, error) {

	tot, _ := r.Getot()
	avg := tot / float64(len(r.dateval))

	return avg, nil
}
