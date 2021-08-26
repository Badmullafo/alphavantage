package server

import (
	_ "encoding/json"
	"fmt"

	_ "github.com/Badmullafo/alphavantage/golang_web/pkg/request"
)

func Startserver() {

	fmt.Printf("Starting server at port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/
}
