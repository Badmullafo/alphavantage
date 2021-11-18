package request

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"encoding/json"
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

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

func GetJson(apiKey, Symbol string, Ndays int) (*Daily, error) {

	//apiKey, Symbol := "RABZYXWVHB8MX5GO", "IBM"
	url := urlS + apiKey + rtype + Symbol
	//init the loc
	loc, _ := time.LoadLocation(tz)

	//set timezone,
	currentTime := time.Now().In(loc)
	fmt.Println("getting data from", url)

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

	d := Daily{}

	fmt.Println("Unmarshalling", currentTime)
	json.Unmarshal([]byte(out), &d)

	return &d, nil

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
