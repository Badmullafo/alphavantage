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

	d := &Daily{}

	fmt.Println("Unmarshalling", currentTime)
	json.Unmarshal([]byte(out), &d)

	return d, nil

}

func (r *Result) getInRange(d *Daily) ([]Dailydata, error) {

	dslice := []Dailydata{}
	loc, _ := time.LoadLocation(tz)

	//set timezone,
	currentTime := time.Now().In(loc)

	for k, v := range d.DD {

		date, err := time.Parse(layoutISO, k)
		if err != nil {
			fmt.Println(err)
		}

		diff := currentTime.Sub(date).Hours()

		//Convert days to hours
		daysH := float64(r.Ndays) * 24

		if diff < daysH {
			dslice = append(dslice, v)

		}
	}
	return dslice, nil
}

func (r *Result) Getot(d []Dailydata, value string) {
	var total float64
	for _, v := range d {
		switch value {
		case "high":
			total = total + v.High
		}
	}
	r.Value = total
	r.Dtype = "total"
}

func (r Result) String() string {
	//var v float64 = r.Value
	return fmt.Sprintf("%.2f", r.Value)
}

/*
func (r *Result) Getavg() {
	r.Getot()
	r.Value = r.Value / float64(len(r.Dateval))
	r.Dtype = "average"

}
*/
