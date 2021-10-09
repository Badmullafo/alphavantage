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

type Result struct {
	Symbol  string                // the stock name .e.g. FORG
	Ndays   int                   // Number of days data to get
	Dtype   string                // data type - total/average
	Dateval map[time.Time]float64 // the date stamp of values returned
	Value   float64
}

func GetJson(apiKey, Symbol string, Ndays int) (*Result, error) {

	//apiKey, Symbol := "RABZYXWVHB8MX5GO", "IBM"
	url := urlS + apiKey + rtype + Symbol
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

	getRes := &Result{Symbol, Ndays, "", make(map[time.Time]float64), 0.0}
	// If there is an error message return here
	if err := gjson.GetBytes(out, "Error Message"); err.String() != "" {
		return nil, errors.New(err.String())
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
		daysH := float64(Ndays) * 24

		if diff < daysH {

			//Inner loop
			value.ForEach(func(key, value gjson.Result) bool {

				if k := key.String(); k == "4. close" {
					v := value.Float()
					//fmt.Printf("The key is:%s, the value is:%f\n", k, v)
					getRes.Dateval[date] = v
				}
				return true // keep iterating
			})
		}
		return true // keep iterating
	})

	return getRes, nil
}

func (r *Result) Getot() {

	var total float64
	for _, value := range r.Dateval {
		total = total + value
	}

	r.Value = total
	r.Dtype = "total"
}

func (r *Result) Getavg() {
	r.Getot()
	r.Value = r.Value / float64(len(r.Dateval))
	r.Dtype = "average"

}
