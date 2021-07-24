package get

import (
	//   "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

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
		Volume     string `json:"6. volume"`
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

	apiKey := "RABZYXWVHB8MX5GO"
	symbol := "IBM"
	url := "https://www.alphavantage.co/query?apikey=" + apiKey + "&function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + symbol

	fmt.Println("Geting data from", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err.Error())
	}

	//fmt.Printf("body = %v", string(body))
	//outputs: {"success":true,"message":"","result":["High":0.43600000,"Low":0.43003737],"Created":"2017-06-25T03:06:46.83"}]}

	var summary = new(Daily)
	err3 := json.Unmarshal(body, &summary)
	if err3 != nil {
		fmt.Println("whoops:", err3)
		//outputs: whoops: <nil>
	}

	//fmt.Printf("\nblah %s \n", summary.DD)

	fmt.Printf("%+v\n", summary.DD)
}
