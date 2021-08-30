package server

import (
	"fmt"
)

func Startserver(value float64) error {

	//fmt.Println("The key is:", key)

	fmt.Printf("The total is %.2f\n", value)

	fmt.Printf("Starting server at port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	*/
	return nil
}
