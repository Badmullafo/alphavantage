package server

import (
	"fmt"
	"github.com/Badmullafo/alphavantage/golang_web/pkg/request"
)

func Startserver() {

	total, err := request.GetJson("RABZYXWVHB8MX5GO", "IBM", 3)

	if err != nil {
		return
	} else {
		fmt.Printf("The total is %.2f\n", total)
	}

	fmt.Printf("Starting server at port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/
}
