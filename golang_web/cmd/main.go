package main

import (
	"fmt"
	"github.com/Badmullafo/alphavantage/golang_web/pkg/request@golang"
	"log"
	"net/http"
)

func main() {

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
