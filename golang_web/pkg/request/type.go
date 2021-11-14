package request

import "time"

type Daily struct {
	MetaData `json:"Meta Data"`
	DD       map[string]Dailydata `json:"Time Series (Daily)"`
}
type MetaData struct {
	Info           string `json:"1. Information"`
	Symbol         string `json:"2. Symbol"`
	Last_refreshed string `json:"3. Last Refreshed"`
	Output_size    string `json:"4. Output Size"`
	Tz             string `json:"5. Time Zone"`
}

type Dailydate struct {
	time.Time
}
type Dailydata struct {
	Open       string `json:"1. open"`
	High       string `json:"2. high"`
	Low        string `json:"3. low"`
	Close      string `json:"4. close"`
	Adj_close  string `json:"5. adjusted close"`
	Volume     string `json:"6. volume"` // `json:"id,string,omitempty"`
	Div_amount string `json:"7. dividend amount"`
	Split_coef string `json:"8. split coefficient"`
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
