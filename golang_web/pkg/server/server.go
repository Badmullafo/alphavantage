package server

import (
	_ "encoding/json"
	"fmt"

	_ "github.com/Badmullafo/alphavantage/golang_web/pkg/request"
)

const (
	layoutISO = "2006-01-02"
	tz        = "US/Eastern"
	layoutUS  = "January 2, 2006"
	urlS      = "https://www.alphavantage.co/query?apikey="
	rtype     = "&function=TIME_SERIES_DAILY_ADJUSTED&symbol="
)

func Startserver() {

	fmt.Printf("Starting server at port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/
}
