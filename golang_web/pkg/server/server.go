package server

import (
	"fmt"
	"github.com/Badmullafo/alphavantage/golang_web/pkg/request"
)

func Startserver(key, stock string, nDays int) {

	fmt.Println("The key is:", key)

	total, err := request.GetJson(key, stock, nDays)

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
