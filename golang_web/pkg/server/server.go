package server

import (
	_ "encoding/json"
	"fmt"

	"github.com/Badmullafo/alphavantage/golang_web/pkg/request"
)

func Startserver() {

	fmt.Println("Scraping data")

	json, err := request.GetJson("RABZYXWVHB8MX5GO", "IBM")

	if err != nil {
		fmt.Println("Something went wrong with the request")
		return
	}

	fmt.Printf("Starting server at port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/
}
