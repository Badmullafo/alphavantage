package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/Badmullafo/golang_web/pkg/get"
)

func main() {

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
