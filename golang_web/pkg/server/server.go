package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type totalHandler struct {
	mu      sync.Mutex // guards n
	message float64
	c       int
}

func (h *totalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.c++
	fmt.Fprintf(w, "the total is %f\n", h.message)
	fmt.Fprintf(w, "the count is %d\n", h.c)
}

func Startserver(path string, value float64) {

	fmt.Printf("Starting server at port 8080\n")

	total := &totalHandler{
		message: value,
	}

	http.Handle(path, total)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
