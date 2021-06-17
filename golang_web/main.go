package main

import (
 //   "bytes"
   // "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {

	apiKey := "RABZYXWVHB8MX5GO"
	symbol := "IBM"
	url :="https://www.alphavantage.co/query?apikey="+apiKey+"&function=TIME_SERIES_DAILY_ADJUSTED&symbol="+symbol

    fmt.Println("Starting the application...")
    response, err := http.Get(url)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }

}