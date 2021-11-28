package request

type Result struct {
	Symbol  string             // the stock name .e.g. FORG
	Ndays   int                // Number of days data to get
	Dtype   string             // data type - total/average
	Dateval map[string]float64 // the date stamp of values returned
	Value   float64
}

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

// NOTE: be careful without the string bit at the end - `json:"1. open,string"` it doesn't know what the original type is and a nil value is initialised
type Dailydata struct {
	Open       float64 `json:"1. open,string"`
	High       float64 `json:"2. high,string"`
	Low        float64 `json:"3. low,string"`
	Close      float64 `json:"4. close,string"`
	Adj_close  float64 `json:"5. adjusted close,string"`
	Volume     uint64  `json:"6. volume,string"`
	Div_amount float64 `json:"7. dividend amount,string"`
	Split_coef float64 `json:"8. split coefficient,string"`
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
