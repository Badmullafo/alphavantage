package request

import (
	//   "bytes"
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	_ "log"
	"net/http"
	"time"
)

type MetaData struct {
	Info           string `json:"1. Information"`
	Symbol         string `json:"2. Symbol"`
	Last_refreshed string `json:"3. Last Refreshed"`
	Output_size    string `json:"4. Output Size"`
	Tz             string `json:"5. Time Zone"`
}

type Dailydata struct {
	Open       string `json:"1. Open"`
	High       string `json:"2. high"`
	Low        string `json:"3. low"`
	Close      string `json:"4. close"`
	Adj_close  string `json:"5. adjusted close"`
	Volume     int64  `json:"6. volume"` // `json:"id,string,omitempty"`
	Div_amount string `json:"7. dividend amount"`
	Split_coef string `json:"8. split coefficient"`
}

type Dailydate struct {
	time.Time
}

type Daily struct {
	MetaData `json:"Meta Data"`
	DD       map[string]Dailydata `json:"Time Series (Daily)"`
}

/* type Daily struct {
   "Time Series (Daily)": {
       "2021-06-17": {
           "1. open": "147.55",
           "2. high": "148.06",
           "3. low": "145.28",
           "4. close": "145.6",
           "5. adjusted close": "145.6",
           "6. volume": "4367387",
           "7. dividend amount": "0.0000",
           "8. split coefficient": "1.0"
       },
*/

func GetJson(url string, jsonobj *Daily) error {

	fmt.Println("Geting data from", url)
	var myClient = &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(jsonobj); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}
