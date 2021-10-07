package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type resultHandler struct {
	mu    sync.Mutex // guards n
	fval  float64
	vtype string
	c     int
}

func (h *resultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.c++
	fmt.Fprintf(w, "The %s is %.2f\n", h.vtype, h.fval)
	fmt.Fprintf(w, "The count is %d\n", h.c)
}

func Startserver(path, vtype string, value float64) {

	fmt.Printf("Starting server at port 8080\n")

	h := &resultHandler{
		vtype: vtype,
		fval:  value,
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.Handle(path, h)
	log.Fatal(s.ListenAndServe())

}
